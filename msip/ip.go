package msip

import (
	"net"
	"strings"
)

const IPV4 = 4

const IPV6 = 6

func GetVersion(val string) int {
	if strings.Contains(val, ":") {
		return IPV4
	}
	return IPV6
}

func IsIPv4(val string) bool {
	ip := net.ParseIP(val)
	return nil != ip && strings.Contains(val, ".")
}

func IsIPv6(val string) bool {
	ip := net.ParseIP(val)
	return nil != ip && strings.Contains(val, ":")
}

func IsIP(val string) bool {
	return IsIPv4(val) || IsIPv6(val)
}

func InRange(subnet string, ipRange string) bool {
	parseIp := func(val string) (net.IP, *net.IPNet, error) {
		if ip := net.ParseIP(val); ip != nil {
			return ip, nil, nil
		}
		return net.ParseCIDR(val)
	}

	ipA, ipnetA, errA := parseIp(subnet)
	ipB, ipnetB, errB := parseIp(ipRange)

	if errA != nil || errB != nil {
		return false
	}

	// 单纯的ip
	if (ipA != nil && ipnetA == nil) && (ipB != nil && ipnetB == nil) {
		return ipA.Equal(ipB)
	}

	// A是ip, B是ip段
	if (ipA != nil && ipnetA == nil) && (ipnetB != nil) {
		return ipnetB.Contains(ipA)
	}

	return false
}
