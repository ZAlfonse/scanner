package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/zalfonse/lumber"
)

var logger *lumber.Logger

func main() {
	blockPtr := flag.String("cidrBlock", "10.0.0.0/8", "CIDR block to scan")
	portPtr := flag.Int("port", 22, "Target port")
	timeoutPtr := flag.Int("timeout", 10, "Time (ms) to wait for established TCP connection")
	loglevelPtr := flag.String("loglevel", "info", "Loglevels: trace, debug, info, quiet, silent")

	flag.Parse()

	switch *loglevelPtr {
	case "trace":
		logger = lumber.NewLogger(lumber.TRACE)
	case "debug":
		logger = lumber.NewLogger(lumber.DEBUG)
	case "quiet":
		logger = lumber.NewLogger(lumber.QUIET)
	case "silent":
		logger = lumber.NewLogger(lumber.SILENT)
	default:
		logger = lumber.NewLogger(lumber.INFO)
	}

	scanCidr(blockPtr, portPtr, timeoutPtr)
}

func scanCidr(block *string, port *int, timeout *int) {
	startIP, ipnet, err := net.ParseCIDR(*block)
	if err != nil {
		logger.Error(err)
	}
	numip := 0
	onlineip := 0
	for ip := startIP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, *port), time.Duration(*timeout)*time.Millisecond)
		if err == nil {
			logger.Success(ip, " online")
			onlineip++
			conn.Close()
		} else {
			logger.Debug(ip, " offline")
		}
		numip++
	}
	logger.Info("Scanned addresses: ", numip)
	logger.Success("Online addresses: ", onlineip)
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
