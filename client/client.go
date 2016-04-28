package main

import (
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/irlndts/altpoint-challenge/thrift/calculator"
)

func handleClient(client *calculator.CalculatorClient) (err error) {
	client.Ping()
	fmt.Println("ping()")

	sum, _ := client.Request("1 + 5 + 3 + 2")
	fmt.Print("1+1=", sum, "\n")

	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	return handleClient(calculator.NewCalculatorClientFactory(transport, protocolFactory))
}
