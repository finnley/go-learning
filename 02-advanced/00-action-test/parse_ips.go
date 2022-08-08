package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func main() {
	//strs, err := parseIpsStr("10.20.30.1-10.20.40.1")
	strs, err := parseFrontEndSipToCIDR("10.20.30.1-10.20.40.1")
	if err != nil {
		return
	}
	fmt.Println(len(strs))
}

const (
	Dash  = "-"
	Slash = "/"
	Comma = ","
)

// ip2int converts an IP to decimal number
func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// int2ip converts decimal number to an IP
func int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func parseFrontEndSipToCIDR(sipInput string) (parsedIPs []string, err error) {
	sipInput = strings.ReplaceAll(sipInput, " ", "")
	switch {
	case strings.Contains(sipInput, Slash) && strings.Contains(sipInput, Dash): // treated as 192.168.10.1/24-192.168.10.10/24
		ipRange := strings.Split(sipInput, Dash)
		if len(ipRange) != 2 {
			return nil, fmt.Errorf("[ERROR] malformed sip[%v], maybe you should use this format: 192.168.10.1/24-192.168.10.10/24", sipInput)
		}
		headIpv4Addr, headIpv4Net, err := net.ParseCIDR(ipRange[0])
		if err != nil {
			return nil, fmt.Errorf("[ERROR] failed to parse CIDR[%v]: %w", ipRange[0], err)
		}
		tailIpv4Addr, tailIpv4Net, err := net.ParseCIDR(ipRange[1])
		if err != nil {
			return nil, fmt.Errorf("[ERROR] failed to parse CIDR[%v]: %w", ipRange[1], err)
		}
		// check network
		if headIpv4Net.String() != tailIpv4Net.String() {
			return nil, fmt.Errorf("[ERROR] your sip range[%v] doesn't belong to the same sub net, the head CIDR belongs to: %v while the tail CIDR belongs to: %v",
				sipInput, headIpv4Net, tailIpv4Net)
		}
		parsedIPs, err = parseIPRange2CIDRs(headIpv4Net, headIpv4Addr.String(), tailIpv4Addr.String())
		if err != nil {
			return nil, fmt.Errorf("[ERROR] failed to parse ip range(from %v to %v): %w", headIpv4Addr, tailIpv4Addr, err)
		}
	case strings.Contains(sipInput, Dash): // treated as 192.168.10.1-192.168.10.10
		ipRange := strings.Split(sipInput, Dash)
		if len(ipRange) != 2 {
			return nil, fmt.Errorf("[ERROR] malformed sip[%v], maybe you should use this format: 192.168.10.1-192.168.10.10", sipInput)
		}
		parsedIPs, err = parseIPRange2CIDRs(nil, ipRange[0], ipRange[1])
		if err != nil {
			return nil, fmt.Errorf("[ERROR] failed to parse ip range(from %v to %v): %w", ipRange[0], ipRange[1], err)
		}
	case strings.Contains(sipInput, Comma): // treated as 192.168.10.1, 192.168.10.9, 192.168.11.1 or 192.168.10.1/24, 192.168.10.9/32, 192.168.11.1/22
		ips := strings.Split(sipInput, Comma)
		for _, ip := range ips {
			if parsedIP, err := parseIP2CIDR(ip); err != nil {
				return nil, fmt.Errorf("[ERROR] failed to parse IP(%v): %w", ip, err)
			} else {
				parsedIPs = append(parsedIPs, parsedIP)
			}
		}
	default: // treated as 192.168.10.1 or 192.168.10.1/24
		if parsedIP, err := parseIP2CIDR(sipInput); err != nil {
			return nil, fmt.Errorf("[ERROR] failed to parse IP(%v): %w", sipInput, err)
		} else {
			return []string{parsedIP}, nil
		}
	}
	return parsedIPs, nil
}

// parseIP2CIDR parsed an CIDR format IP(10.186.12.1/24) or an normal IP(10.186.12.1) to CIDR
func parseIP2CIDR(IP string) (string, error) {
	ipNet := &net.IPNet{}
	if strings.Contains(IP, Slash) {
		headIpv4Addr, headIpv4Net, err := net.ParseCIDR(IP)
		if err != nil {
			return "", fmt.Errorf("[ERROR] failed to parse CIDR[%v]: %w", IP, err)
		}
		ipNet.IP = headIpv4Addr
		ipNet.Mask = headIpv4Net.Mask
	} else {
		parsedIP := net.ParseIP(IP)
		if parsedIP == nil {
			return "", fmt.Errorf("[ERROR] invalid IP: %v", IP)
		}
		ipNet.IP = parsedIP
		ipNet.Mask = []byte{255, 255, 255, 255}
	}
	return ipNet.String(), nil
}

func parseIPRange2CIDRs(ipNet *net.IPNet, headIPString, tailIPString string) ([]string, error) {
	headIP := net.ParseIP(headIPString)
	if headIP == nil {
		return nil, fmt.Errorf("[ERROR] headIPString(%v) is invalid", headIPString)
	}
	tailIP := net.ParseIP(tailIPString)
	if tailIP == nil {
		return nil, fmt.Errorf("[ERROR] tailIPString(%v) is invalid", tailIPString)
	}
	if bytes.Compare(tailIP, headIP) < 0 {
		return nil, fmt.Errorf("[ERROR] tailIPString(%v) should be bigger than headIPString(%v)", tailIPString, headIPString)
	}
	if ipNet == nil {
		ipNet = &net.IPNet{}
	}
	if ipNet.Mask == nil {
		ipNet.Mask = []byte{255, 255, 255, 255}
	}

	head := ip2int(headIP)
	tail := ip2int(tailIP)
	var n *net.IPNet
	var parsedIPs []string
	for ; head <= tail; head++ {
		ip := int2ip(head)
		n = &net.IPNet{
			Mask: ipNet.Mask,
			IP:   ip,
		}
		parsedIPs = append(parsedIPs, n.String())
	}
	return parsedIPs, nil
}

func parseIpsStr(ipsStr string) ([]string, error) {
	var sips []string
	for _, sipStr := range strings.Split(ipsStr, ",") {
		if "" == sipStr {
			continue
		}
		ips, err := parseIpStr2(sipStr)
		if nil != err {
			return nil, err
		}
		sips = append(sips, ips...)
	}
	return sips, nil
}

func parseIpStr2(ipStr string) ([]string, error) {
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
