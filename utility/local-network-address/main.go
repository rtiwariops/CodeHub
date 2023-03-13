package main

import (
    "fmt"
    "net"
)

func main() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        panic(err)
    }
    for _, addr := range addrs {
        ipNet, ok := addr.(*net.IPNet)
        if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
            fmt.Println(ipNet.IP)
        }
    }
}

