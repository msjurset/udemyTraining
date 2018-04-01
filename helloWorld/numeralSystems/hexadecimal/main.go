package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("Dec: %d\nBin:%b\nHex:%#x\n", i, i, i)
		//fmt.Printf("Dec: %d\nBin:%b\nHex:%#X\n", i, i, i)
	}

}
