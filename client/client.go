package main

import (
	"bufio"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/irlndts/altpoint-challenge/thrift/calculator"
	"io"
	"os"
)

func handleClient(client *calculator.CalculatorClient) (err error) {
	// interactive input
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println()
			os.Exit(0)
		}

		switch string(s) {
		case "quit\n":
			fmt.Println("Chao")
			return nil
		case "ping\n":
			client.Ping()
			fmt.Println("pong()")
			break
		case "\n":
			break
		default:
			sum, err := client.Request(s)
			if err != nil {
				fmt.Println(err)
				fmt.Println("####################")
				fmt.Println("Usage: ")
				fmt.Println("- quit - to leave the program")
				fmt.Println("- ping - to check if servers is still available")
				fmt.Println("- math expression WITH SPACES between elements.")
				break
			}
			fmt.Println("Result:", sum)
			break
		}
	}

	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(addr)

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
