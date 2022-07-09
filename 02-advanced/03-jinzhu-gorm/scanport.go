package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	//host := "111.231.87.78"
	host := "127.0.0.1"
	port := 3380
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 30*time.Second)
	//tcpAddr := net.TCPAddr{
	//	IP:   net.ParseIP(host),
	//	Port: port,
	//}
	//conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	fmt.Println(err)

	if err == nil {
		fmt.Println("success")
		conn.Close()
		conn = nil
	} else {
		fmt.Println("failed")
	}
}

func main4() {
	ips := "172.20.134.7/20"
	hostList, err := ParseIpsStr(ips)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hostList)
}

func ParseIpsStr(ipsStr string) ([]string, error) {
	var sips []string
	for _, sipStr := range strings.Split(ipsStr, ",") {
		if "" == sipStr {
			continue
		}
		ips, err := parseIpStr(sipStr)
		if nil != err {
			return nil, err
		}
		sips = append(sips, ips...)
	}
	return sips, nil
}

func parseIpStr(ipStr string) ([]string, error) {
	inc := func(ip net.IP) {
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] > 0 {
				break
			}
		}
	}
	ipStr = strings.Replace(ipStr, "\n", "", -1)
	ipStr = strings.Replace(ipStr, " ", "", -1)
	if strings.Contains(ipStr, "-") {
		ips := strings.Split(ipStr, "-")
		if len(ips) != 2 {
			return nil, fmt.Errorf("\"%v\" is of invalid format", ipStr)
		}
		ip1 := net.ParseIP(ips[0])
		if nil == ip1 {
			return nil, fmt.Errorf("\"%v\" is of invalid format", ips[0])
		}
		ip2 := net.ParseIP(ips[1])
		if nil == ip2 {
			return nil, fmt.Errorf("\"%v\" is of invalid format", ips[1])
		}
		var ret []string
		for ip := ip1; !ip.Equal(ip2); inc(ip) {
			ret = append(ret, ip.String())
			fmt.Println(ret)
			if len(ret) > 1024 {
				return nil, fmt.Errorf("cannot add more than 1024 ips at once")
			}
		}
		return append(ret, ip2.String()), nil
	} else if strings.Contains(ipStr, "/") {
		//CIDR
		ip, ipnet, err := net.ParseCIDR(ipStr)
		if err != nil {
			return nil, err
		}

		var ips []string
		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			ips = append(ips, ip.String())
			fmt.Println(ips)
			if len(ips) > 1024 {
				return nil, fmt.Errorf("cannot add more than 1024 ips at once")
			}
		}
		// remove network address and broadcast address
		return ips[1 : len(ips)-1], nil
	} else {
		ip := net.ParseIP(ipStr)
		if nil == ip {
			return nil, fmt.Errorf("\"%v\" is of invalid format", ipStr)
		}
		return []string{ip.String()}, nil
	}
}
