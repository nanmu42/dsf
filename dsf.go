package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

var (
	root = flag.String("root", "./", "root of files")
	port = flag.Int("port", 8080, "port to listen")
)

func main() {
	flag.Parse()

	var err error
	defer func() {
		if err != nil {
			fmt.Printf("fatal error: %s\n", err)
		}
	}()

	interfaces, err := net.Interfaces()
	if err != nil {
		err = fmt.Errorf("listing net interfaces: %w", err)
		return
	}
	for _, i := range interfaces {
		var addrs []net.Addr
		addrs, err = i.Addrs()
		if err != nil {
			err = fmt.Errorf("i.Addrs: %w", err)
			return
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			fmt.Printf("listening on: http://%s:%d\n", ip, *port)
		}
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(*root)))
	if err != nil {
		err = fmt.Errorf("http.ListenAndServe: %w", err)
		return
	}
}
