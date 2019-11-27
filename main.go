package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/netutil"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {

	var (
		host     string
		port_str string
	)

	flag.StringVar(&host, "host", "0.0.0.0", "host")
	flag.StringVar(&port_str, "p", "80", "port")
	flag.Parse()

	port, err := strconv.ParseUint(port_str, 10, 16)

	if err != nil {
		log.Fatalln(err)
	}

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		log.Fatalln(err)
	}

	limit_ln := netutil.LimitListener(ln, 100)
	defer limit_ln.Close()

	server := &http.Server{
		Handler: GenerateRouter(),
	}

	err = server.Serve(limit_ln)

	if err != nil {
		log.Fatalln(err)
	}
}
