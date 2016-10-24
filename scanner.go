package main

import (
	"fmt"
	"log"
	"net"
  "flag"
  "time"
)

func main() {
  blockPtr := flag.String("cidrBlock", "10.0.0.0/8", "CIDR block to scan")
  portPtr := flag.Int("port", 23, "Target port")
  timeoutPtr := flag.Int("timeout", 10, "Time (ms) to wait for established TCP connection")
  verbosePtr := flag.Bool("v", false, "Print everything")

  flag.Parse()

  scanCidr(blockPtr, portPtr, timeoutPtr, verbosePtr)
}

func scanCidr(block *string, port *int, timeout *int, verbose *bool) {
  ip, ipnet, err := net.ParseCIDR(*block)
  if err != nil {
    log.Fatal(err)
  }
  numip := 0
  onlineip := 0
  for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
  	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, *port), time.Duration(*timeout) * time.Millisecond)
    if err == nil {
      fmt.Println(ip, "online")
      onlineip += 1
      conn.Close()
    } else {
      if(*verbose){
        fmt.Println(ip, "offline")
      }
    }
    numip += 1
  }

  fmt.Println("Scanned addresses:", numip)
  fmt.Println("Online addresses:", onlineip)
}

func inc(ip net.IP) {
  for j := len(ip)-1; j>=0; j-- {
    ip[j]++
    if ip[j] > 0 {
      break
    }
  }
}
