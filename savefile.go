package psedit

import (
	"bytes"
	"encoding/binary"
	"github.com/lunixbochs/struc"
	"io"
)

type Item uint8

const (
	Nothing Item = 0
	MaxItem Item = 64
)

// A Phantasy Star 1 save file consists of five saved games.
// A header of size 0x200 exists at position 0x100 in the save file.
// Assuming saved game indexes ("GameIndex") numbered 0 through 4, each saved
// game can be found in the file at location `0x500 + 0x400 * gamenum`.
// In other words, saved games start at index 0x500 and are each of size 0x400.
// Each saved game consists of four packed 16-byte PlayerRecord structures,
// representing Alis, Myau, Odin, and Noah. At offset 0xC0 in each saved game,
// there is a 32-byte array of inventory items. At offset 0xE0 in the saved
// game, a two-byte value can be found representing the number of meseta
// followed by a one-byte value representing the number of inventory items.

type Header [0x200]uint8

var NameOffsets = [...]uint16{
	0x2a,
	0x4e,
	0x72,
	0x96,
	0xba,
}

var DeletedFlagOffsets = [...]uint16{
	0x101,
	0x102,
	0x103,
	0x104,
	0x105,
}

// PlayerRecord is a 16-byte (0x10) structure representing each character.
type PlayerRecord struct {
	Alive              bool
	CurrentHP          uint8
	CurrentMP          uint8
	Experience         uint16 `struc:"little"`
	Level              uint8
	MaxHP              uint8
	MaxMP              uint8
	Attack             uint8
	Defense            uint8
	Weapon             Item
	Armor              Item
	Shield             Item
	State              uint8
	NumCombatSpells    uint8
	NumNonCombatSpells uint8
}

// Pack returns a bytes.Buffer object suitable for writing to a save file.
// (Using the Go structure directly results in too much padding.)
func (record *PlayerRecord) Pack() bytes.Buffer {
	var buffer = bytes.Buffer{}
	err := struc.Pack(&buffer, record)
	if err != nil {
		panic(err)
	}
	return buffer
}

type Inventory [32]Item

type SavedGame struct {
	// 0x40 bytes of character date
	Characters [4]PlayerRecord
	// Pad out to offset 0xC0
	Padding1          [0x80]uint8
	Inventory         Inventory
	Meseta            uint16 `struc:"little"`
	NumInventoryItems uint8
	Padding2          [0x31d]uint8
}

func ReadSavedGame(r io.Reader, savedGame *SavedGame) error {
	var err error
	for i := 0; i < len(savedGame.Characters); i++ {
		err := struc.Unpack(r, &savedGame.Characters[i])
		if err != nil {
			return err
		}
	}
	err = binary.Read(r, binary.BigEndian, &savedGame.Padding1)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.BigEndian, &savedGame.Inventory)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &savedGame.Meseta)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.BigEndian, &savedGame.NumInventoryItems)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.BigEndian, &savedGame.Padding2)
	if err != nil {
		return err
	}
	return nil
}

// Pack returns a bytes.Buffer object suitable for writing to a save file.
// (Using the Go structure directly results in too much padding.)
func (record *SavedGame) Pack() bytes.Buffer {
	var buffer = bytes.Buffer{}
	err := struc.Pack(&buffer, record)
	if err != nil {
		panic(err)
	}
	return buffer
}

var ExpectedMagic = []uint8("PHANTASY STAR         BACKUP RAM" +
	"PROGRAMMED BY          NAKA YUJI")

type SaveFile struct {
	Magic  [0x100]uint8
	Header [0x200]uint8
	// Pad out to 0x500
	Padding1 [0x200]uint8
	// Saved games (starting at offset 0x500)
	Games [5]SavedGame
	// Pad out to 16 kilobytes
	Padding2 [0x2700]uint8
}

// Pack returns a bytes.Buffer object suitable for writing to a save file.
// (Using the Go structure directly results in too much padding.)
func (record *SaveFile) Pack() bytes.Buffer {
	var buffer = bytes.Buffer{}
	err := struc.Pack(&buffer, record)
	if err != nil {
		panic(err)
	}
	return buffer
}

func (record *SaveFile) HasValidMagic() bool {
	if bytes.HasPrefix(record.Magic[:], ExpectedMagic) {
		return true
	}
	return false
}

func ReadSaveFile(r io.Reader) (*SaveFile, error) {
	saveFile := &SaveFile{}
	var err error
	err = binary.Read(r, binary.BigEndian, &saveFile.Magic)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &saveFile.Header)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &saveFile.Padding1)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(saveFile.Games); i++ {
		err = ReadSavedGame(r, &saveFile.Games[i])
		if err != nil {
			return nil, err
		}
	}
	err = binary.Read(r, binary.BigEndian, &saveFile.Padding2)
	return saveFile, nil
}

var Items = map[Item]string{
	Nothing: "[nothing]",
	1:       "Wooden Cane",
	2:       "Short Sword",
	3:       "Iron Sword",
	4:       "Wand",
	5:       "Iron Fang",
	6:       "Iron Axe",
	7:       "Titanium Sword",
	8:       "Ceramic Sword",
	9:       "Needle Gun",
	10:      "Silver Fang",
	11:      "Heat Gun",
	12:      "Light Saber",
	13:      "Laser Gun",
	14:      "Laconium Sword",
	15:      "Laconium Axe",
	16:      "Lithium Armor",
	17:      "White Mantle",
	18:      "Light Suit",
	19:      "Iron Armor",
	20:      "Thick Fur",
	21:      "Zirconium Armor",
	22:      "Diamond Armor",
	23:      "Laconium Armor",
	24:      "Fraid Mantle",
	25:      "Lithium Shield",
	26:      "Bronze Shield",
	27:      "Iron Shield",
	28:      "Ceramic Shield",
	29:      "Glove",
	30:      "Laser Shield",
	31:      "Mirror Shield",
	32:      "Lacomium Shield",
	33:      "Land Rover",
	34:      "Hovercraft",
	35:      "Ice Digger",
	36:      "Cola",
	37:      "Burger",
	38:      "Flute",
	39:      "Flash",
	40:      "Escaper",
	41:      "Transfer",
	42:      "Magic Hat",
	43:      "Alsulin",
	44:      "Polymtrl",
	45:      "Dungeon Key",
	46:      "Sphere",
	47:      "Torch",
	48:      "Prism",
	49:      "Nuts",
	50:      "Hapsby",
	51:      "Road Pass",
	52:      "Passport",
	53:      "Compass",
	54:      "Cake",
	55:      "Letter",
	56:      "Laconium Pot",
	57:      "Magic Lamp",
	58:      "Amber Eye",
	59:      "Gas Shield",
	60:      "Crystal",
	// The following items are not usable.
	61: "M System",
	62: "Miricle Key",
	63: "Zillion",
	64: "Secrets",
}

// String returns a string representation for the specified Item.
func (item Item) String() string {
	return Items[item]
}
