package main

import (
	"fmt"
	"github.com/irlndts/go-rpn"
)

type CalculatorHandler struct {
}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{}
}

// Ping request checks connection to the server
func (p *CalculatorHandler) Ping() (err error) {
	// TODO check status of the server etc
	fmt.Println("Ping request")
	return nil
}

// Request of some math expression
// Input: expresson as string like "1 + 5 - 6 * 9"
// Returnis result of calculation as float64 or error
func (p *CalculatorHandler) Request(expr string) (r float64, err error) {
	fmt.Println("Request (", expr, ")")
	res, err := rpn.Calc(rpn.Parse(expr))

	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}

	fmt.Println("Response: ", res)

	return res, nil
}
