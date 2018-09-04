package psedit

import (
	"io/ioutil"
	"os"
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

func TestSaveFileHasGoodLength(t *testing.T) {
	data, err := ioutil.ReadFile("data/phanstar.sav")
	if err != nil {
		panic(err)
	}
	if len(data) != 0x4000 {
		t.Errorf("Expected 0x4000 bytes; save file was 0x%04x bytes", len(data))
	}
}

func TestReadSaveFile(t *testing.T) {
	saveFile := &SaveFile{}
	f, err := os.Open("data/phanstar.sav")
	if err != nil {
		panic(err)
	}
	saveFile, err = ReadSaveFile(f)
	if !saveFile.HasValidMagic() {
		t.Error("Invalid save file.")
	}
}

func TestCheckMagicFailsForBadMagic(t *testing.T) {
	saveFile := &SaveFile{}
	f, err := os.Open("data/invalid_magic.sav")
	if err != nil {
		panic(err)
	}
	saveFile, err = ReadSaveFile(f)
	if saveFile.HasValidMagic() {
		t.Error("Valid magic in save file. (Expected bad magic.)")
	}
}

func TestExperiments(t *testing.T) {
	var item Item = 23
	t.Log(item.String())
	var playerRecord = PlayerRecord{}
	var saveFile = SaveFile{}
	t.Log(unsafe.Sizeof(playerRecord))
	var buffer = playerRecord.Pack()
	t.Log(buffer.Len())
	t.Log(unsafe.Sizeof(saveFile))
	buffer = saveFile.Pack()
	t.Logf("0x%x\n", buffer.Len())
}
