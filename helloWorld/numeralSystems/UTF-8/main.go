package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("Dec: %d\tBin:%b\tHex:%#X\tUTF-8:%q\n", i, i, i, i)
		//fmt.Printf("Dec: %d\tBin:%b\tHex:%#x\tUTF-8:%q\n", i, i, i,i)

	}
}
