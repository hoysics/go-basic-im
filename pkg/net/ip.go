package net

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() (ip string, err error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range as {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

func GetLocalServerHost(addrOrPort string) (host string, err error) {
	ip, err := GetLocalIP()
	if err != nil {
		return
	}
	if strings.Contains(addrOrPort, ":") {
		ss := strings.Split(addrOrPort, ":")
		if len(ss) != 2 {
			return "", errors.New(`非法的参数格式`)
		}
		return fmt.Sprintf("%s:%s", ip, ss[1]), nil
	}
	return fmt.Sprintf("%s:%s", ip, addrOrPort), nil
}
