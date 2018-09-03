package main

import (
	"fmt"
	"github.com/mpontillo/psedit"
	"unsafe"
)

func main() {
	fmt.Println("Testing.")
	fmt.Println(psedit.Items[23])
	var playerRecord = psedit.PlayerRecord{}
	fmt.Println(unsafe.Sizeof(playerRecord))
	var buffer = playerRecord.Pack()
	fmt.Println(buffer.Len())
	fmt.Println(len(psedit.Inventory{}))
}
