package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	blockPtr := flag.String("cidrBlock", "10.0.0.0/8", "CIDR block to scan")
	portPtr := flag.Int("port", 22, "Target port")
	timeoutPtr := flag.Int("timeout", 10, "Time (ms) to wait for established TCP connection")
	debugPtr := flag.Bool("v", false, "Extra debug info")
	tracePtr := flag.Bool("vv", false, "Print everything")
	quietPtr := flag.Bool("q", false, "Be quieter")
	silentPtr := flag.Bool("qq", false, "Be silent")

	flag.Parse()

	if *tracePtr {
		initLogger(TRACE)
	} else if *debugPtr {
		initLogger(DEBUG)
	} else if *quietPtr {
		initLogger(QUIET)
	} else if *silentPtr {
		initLogger(SILENT)
	} else {
		initLogger(INFO)
	}

	scanCidr(blockPtr, portPtr, timeoutPtr)
}

func scanCidr(block *string, port *int, timeout *int) {
	startIP, ipnet, err := net.ParseCIDR(*block)
	if err != nil {
		errorLog.Println(err)
	}
	numip := 0
	onlineip := 0
	for ip := startIP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, *port), time.Duration(*timeout)*time.Millisecond)
		if err == nil {
			successLog.Println(ip, " online")
			onlineip++
			conn.Close()
		} else {
			debugLog.Println(ip, " offline")
		}
		numip++
	}
	infoLog.Println("Scanned addresses: ", numip)
	successLog.Println("Online addresses: ", onlineip)
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
