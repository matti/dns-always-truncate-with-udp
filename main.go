package main

import (
	"fmt"
	"log"
	"os"

	"github.com/miekg/dns"
)

func handleDnsRequest(w dns.ResponseWriter, request *dns.Msg) {
	fmt.Println("handle", request.Question)
	c := dns.Client{
		Net: "udp",
	}

	response, _, _ := c.Exchange(request, "8.8.8.8:53")
	if server.Net == "udp" {
		response.Truncated = true
		response.Answer = []dns.RR{}
	}

	response.SetReply(request)
	w.WriteMsg(response)
}

var server dns.Server

func main() {
	net := os.Args[1]

	dns.HandleFunc(".", handleDnsRequest)

	server = dns.Server{
		Addr: ":53",
		Net:  net,
	}

	defer server.Shutdown()

	log.Println("Starting at :53, using ", net)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
