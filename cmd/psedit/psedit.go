package main

import (
	"fmt"
	"github.com/mpontillo/psedit"
	"unsafe"
)

func main() {
	var item psedit.Item = 23
	fmt.Println(item.String())
	var playerRecord = psedit.PlayerRecord{}
	var saveFile = psedit.SaveFile{}
	fmt.Println(unsafe.Sizeof(playerRecord))
	var buffer = playerRecord.Pack()
	fmt.Println(buffer.Len())
	fmt.Println(len(psedit.Inventory{}))
	fmt.Println(unsafe.Sizeof(saveFile))
	buffer = saveFile.Pack()
	fmt.Printf("0x%x\n", buffer.Len())
}
