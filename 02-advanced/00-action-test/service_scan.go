package main

import (
	"fmt"
	"go-learning/02-advanced/vscan"
	"net"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var hosts = "127.0.0.1"
var ports = "3306,3357,6379,6380"
var timeout int
var model = "tcp"
var outFile = "a.txt"

func main() {
	var hostLists []string
	hostLists = append(hostLists, hosts)

	var AliveHosts []string
	var AliveAddress []string
	//var TagetBanners []string

	AliveHosts, AliveAddress = TCPportScan(hostLists, ports, model, timeout)
	for _, host := range AliveHosts {
		fmt.Printf("(TCP) Target '%s' is alive\n", host)
	}
	for _, addr := range AliveAddress {
		fmt.Println(addr)
	}

	if len(AliveAddress) > 0 {
		//TagetBanners = vscan.GetProbes(AliveAddress)
		fmt.Println(AliveAddress)
		vscan.GetProbes(AliveAddress)

	}
	//fmt.Println(TagetBanners)
	//for _, taget := range TagetBanners {
	//	fmt.Println(taget)
	//}

	//if outFile != "" && pathCheck(outFile) == true && len(AliveHosts) != 0 {
	//	f, _ := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	//	for _, host := range AliveHosts {
	//		f.WriteString(host + "\n")
	//	}
	//	for _, addr := range AliveAddress {
	//		f.WriteString(addr + "\n")
	//	}
	//	for _, taget := range TagetBanners {
	//		f.WriteString(taget + "\n")
	//	}
	//	fmt.Printf("Output the scanning information in %s\n", outFile)
	//	defer f.Close()
	//}
}

func TCPportScan(hostslist []string, ports string, model string, timeout int) ([]string, []string) {
	var AliveAddress []string
	var aliveHosts []string
	probePorts := parsePort(ports)
	lm := 20
	if len(hostslist) > 5 && len(hostslist) <= 50 {
		lm = 40
	} else if len(hostslist) > 50 && len(hostslist) <= 100 {
		lm = 50
	} else if len(hostslist) > 100 && len(hostslist) <= 150 {
		lm = 60
	} else if len(hostslist) > 150 && len(hostslist) <= 200 {
		lm = 70
	} else if len(hostslist) > 200 {
		lm = 75
	}

	thread := 5
	if len(probePorts) > 500 && len(probePorts) <= 4000 {
		thread = len(probePorts) / 100
	} else if len(probePorts) > 4000 && len(probePorts) <= 6000 {
		thread = len(probePorts) / 200
	} else if len(probePorts) > 6000 && len(probePorts) <= 10000 {
		thread = len(probePorts) / 350
	} else if len(probePorts) > 10000 && len(probePorts) < 50000 {
		thread = len(probePorts) / 400
	} else if len(probePorts) >= 50000 && len(probePorts) <= 65535 {
		thread = len(probePorts) / 500
	}

	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	limiter := make(chan struct{}, lm)
	aliveHost := make(chan string, lm/2)
	go func() {
		for s := range aliveHost {
			fmt.Println(s)
		}
	}()
	for _, host := range hostslist {
		wg.Add(1)
		limiter <- struct{}{}
		go func(host string) {
			defer wg.Done()
			if aliveAdd, err := ScanAllports(host, probePorts, thread, 5*time.Second, model, timeout); err == nil && len(aliveAdd) > 0 {
				mutex.Lock()
				aliveHosts = append(aliveHosts, host)
				for _, addr := range aliveAdd {
					AliveAddress = append(AliveAddress, addr)
				}
				mutex.Unlock()
			}
			<-limiter
		}(host)
	}
	wg.Wait()
	close(aliveHost)
	return aliveHosts, AliveAddress
}

func ScanAllports(address string, probePorts []int, threads int, timeout time.Duration, model string, adjustedTimeout int) ([]string, error) {
	ports := make(chan int, 20)
	results := make(chan string, 10)
	done := make(chan bool, threads)

	for worker := 0; worker < threads; worker++ {
		go ProbeHosts(address, ports, results, done, model, adjustedTimeout)
	}

	for _, port := range probePorts {
		ports <- port
	}
	close(ports)

	var responses = []string{}
	for {
		select {
		case found := <-results:
			responses = append(responses, found)
		case <-done:
			threads--
			if threads == 0 {
				return responses, nil
			}
		case <-time.After(timeout):
			return responses, nil
		}
	}
}

func ProbeHosts(host string, ports <-chan int, respondingHosts chan<- string, done chan<- bool, model string, adjustedTimeout int) {
	Timeout := time.Duration(adjustedTimeout) * time.Second
	for port := range ports {
		start := time.Now()
		con, err := net.DialTimeout("tcp4", fmt.Sprintf("%s:%d", host, port), time.Duration(adjustedTimeout)*time.Second)
		duration := time.Now().Sub(start)
		if err == nil {
			defer con.Close()
			address := host + ":" + strconv.Itoa(port)
			if model == "tcp" {
				fmt.Printf("(TCP) Target %s is open\n", address)
			} else {
				fmt.Println(address)
			}
			respondingHosts <- address
		}
		if duration < Timeout {
			difference := Timeout - duration
			Timeout = Timeout - (difference / 2)
		}
	}
	done <- true
}

func parsePort(ports string) []int {
	var scanPorts []int
	slices := strings.Split(ports, ",")
	for _, port := range slices {
		port = strings.Trim(port, " ")
		upper := port
		if strings.Contains(port, "-") {
			ranges := strings.Split(port, "-")
			if len(ranges) < 2 {
				continue
			}
			sort.Strings(ranges)
			port = ranges[0]
			upper = ranges[1]
		}
		start, _ := strconv.Atoi(port)
		end, _ := strconv.Atoi(upper)
		for i := start; i <= end; i++ {
			scanPorts = append(scanPorts, i)
		}
	}
	return scanPorts
}

func pathCheck(files string) bool {
	path, _ := filepath.Split(files)
	_, err := os.Stat(path)
	if err == nil {
		_, err2 := os.Stat(files)
		if err2 == nil {
			return false
		}
		if os.IsNotExist(err2) {
			return true
		}
	} else {
		err3 := os.MkdirAll(path, os.ModePerm)
		if err3 == nil {
			return true
		} else {
			return false
		}
	}
	return false
}
