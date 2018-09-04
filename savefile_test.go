package psedit

import (
	"bytes"
	"fmt"
	"github.com/lunixbochs/struc"
	"io/ioutil"
	"testing"
	"unsafe"
)

func TestSaveFilePacksCorrectly(t *testing.T) {
	var saveFile = SaveFile{}
	var buffer = saveFile.Pack()
	var length = buffer.Len()
	if length != 0x4000 {
		t.Errorf("Expected 0x4000 bytes; save file was 0x%04x bytes", length)
	}
}

func TestExperiments(t *testing.T) {
	var item Item = 23
	fmt.Println(item.String())
	var playerRecord = PlayerRecord{}
	var saveFile = SaveFile{}
	fmt.Println(unsafe.Sizeof(playerRecord))
	var buffer = playerRecord.Pack()
	fmt.Println(buffer.Len())
	fmt.Println(unsafe.Sizeof(saveFile))
	buffer = saveFile.Pack()
	fmt.Printf("0x%x\n", buffer.Len())
}

func TestReadFile(t *testing.T) {
	saveFile := &SaveFile{}
	data, err := ioutil.ReadFile("data/phanstar.sav")
	if err != nil {
		panic(err)
	}
	if len(data) != 0x4000 {
		t.Errorf("Expected 0x4000 bytes; save file was 0x%04x bytes", len(data))
	}
	/*
		f, err := os.Open("data/phanstar.sav")
		if err != nil {
			panic(err)
		}
	*/
	buffer := bytes.NewBuffer(data)
	struc.Unpack(buffer, saveFile)
	t.Log(saveFile)
}
