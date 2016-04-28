package main

import (
	"fmt"
	"github.com/irlndts/go-rpn"
	//"github.com/irlndts/altpoint-challenge/thrift/calculator"
	//"strconv"
)

type CalculatorHandler struct {
}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{}
}

func (p *CalculatorHandler) Ping() (err error) {
	fmt.Print("ping()\n")
	return nil
}

func (p *CalculatorHandler) Request(expr string) (r float64, err error) {
	fmt.Print("Request (", expr, ")\n")
	res, err := rpn.Calc(rpn.Parse(expr))
	if err != nil {
		return 0, err
	}
	fmt.Println(res)

	return res, nil
}
