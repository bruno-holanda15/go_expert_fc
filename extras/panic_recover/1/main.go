package main

import (
	"fmt"
)

func uauPanic() {
	panic("Something uau wrong")
}

func topPanic() {
	panic("Something top wrong")

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "Something uau wrong" {
				fmt.Println("Recovered uau")
			}

			if r == "Something top wrong" {
				fmt.Println("Recovered top")
			}
		}
	}()

	uauPanic()
	topPanic()
}
