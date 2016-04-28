package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
)

func Usage() {
	fmt.Println("Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Println()
}

func main() {
	flag.Usage = Usage
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		Usage()
		os.Exit(1)
	}

	var protocolFactory thrift.TProtocolFactory
	// set binary protocol. There are some other protocol types like json etc, but we don't need it at the moment
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	if err := runClient(transportFactory, protocolFactory, *addr); err != nil {
		fmt.Println("error running client:", err)
	}
}
