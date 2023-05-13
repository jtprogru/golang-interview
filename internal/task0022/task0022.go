/*
https://www.codewars.com/kata/526989a41034285187000de4/train/go
*/

package task0022

import (
	"net"
	"strconv"
	"strings"
)

func Solution(start string, end string) int {
	startIP := net.ParseIP(start).To4()
	endIP := net.ParseIP(end).To4()
	startInt := ipToInt(startIP)
	endInt := ipToInt(endIP)

	return endInt - startInt
}

func ipToInt(ip net.IP) int {
	ipParts := strings.Split(ip.String(), ".")
	res := 0
	for _, ipPart := range ipParts {
		num, _ := strconv.Atoi(ipPart)
		res = res*256 + num
	}
	return res
}
