package main

import (
	"fmt"
	"go_fibo/cmd/app"
)

func main() {

	Fiby(3)
}

func Fiby(n int) {

	iterRes := app.FibonacciIteration(n)

	curRes := app.FibonacciRecursive(n)

	dpRes := app.FibonacciDP(n)

	fmt.Printf("Iterate Fibanacci:%v \n",iterRes)

	fmt.Printf("Recursive Fibanacci:%v \n",curRes)

	fmt.Printf("Dynamic Programming Algorithm:%v \n",dpRes)

}
