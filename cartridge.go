package main

import (
	"encoding/binary"
	"os"
)

type Mirror int

const (
	Horizontal Mirror = iota
	Vertical
	OnScreenHi
	OnScreenLo
)

type CartridgeHeader struct {
	Name         [4]byte
	PrgROMChunks uint8
	ChrROMChunks uint8
	Mapper1      uint8
	Mapper2      uint8
	PrgRAMSize   uint8
	TVSystem1    uint8
	TVSystem2    uint8
	Unused       [5]byte
}

type Cartridge struct {
	imageValid bool
	mirror     Mirror
	header     CartridgeHeader

	mapperId uint8 // Mapper ID
	prgBanks uint8 // PRG banks
	chrBanks uint8 // CHR banks

	// PrgMemory is a vector of uint8s
	prgMemory []byte
	// ChrMemory is a vector of uint8s
	chrMemory []byte

	mapper Mapper
}

func NewCartridge(filename string) *Cartridge {
	var mapper Mapper
	var imageValid bool
	var prgBanks uint8
	var prgMemory []byte
	var chrBanks uint8
	var chrMemory []byte
	var header CartridgeHeader

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close() // Close when we're done

	// Read the header from the file
	headerBytes := make([]byte, binary.Size(header))
	_, err = file.Read(headerBytes)
	if err != nil {
		panic(err)
	}

	// If a "trainer" exists, skip it
	if headerBytes[6] != 0 {
		_, err = file.Seek(512, 1)
		if err != nil {
			panic(err)
		}
	}

	// Determine mapper id
	mapperId := ((headerBytes[6] >> 4) << 4) | (headerBytes[7] >> 4)
	mirror := Horizontal
	if headerBytes[6]&1 == 1 {
		mirror = Vertical
	}

	fileType := 1

	if fileType == 0 {
	} else if fileType == 1 {
		// Populate PRG banks and allocate memory
		prgBanks = headerBytes[4]
		prgMemory = make([]byte, int(prgBanks)*16384)
		_, err = file.Read(prgMemory)
		if err != nil {
			panic(err)
		}

		// Populate CHR banks and allocate memory
		chrBanks = headerBytes[5]
		chrMemory = make([]byte, int(chrBanks)*8192)
		_, err = file.Read(chrMemory)
		if err != nil {
			panic(err)
		}
	}

	// Load the appropriate mapper
	switch mapperId {
	case 0:
		mapper = &Mapper000{}
		imageValid = true
	}

	// Return a new cartridge
	c := &Cartridge{
		mirror:     mirror,
		mapperId:   mapperId,
		prgBanks:   prgBanks,
		chrBanks:   chrBanks,
		prgMemory:  prgMemory,
		chrMemory:  chrMemory,
		mapper:     mapper,
		imageValid: imageValid,
	}

	return c
}

func (c *Cartridge) ImageValid() bool {
	return c.imageValid
}

func (c *Cartridge) cpuRead(addr uint16, data *byte) bool {
	mappedAddress := uint32(0)
	if c.mapper.cpuMapRead(addr, mappedAddress) {
		*data = c.prgMemory[mappedAddress]
		return true
	} else {
		return false
	}
}

func (c *Cartridge) cpuWrite(addr uint16, data byte) bool {
	mappedAddress := uint32(0)
	if c.mapper.cpuMapWrite(addr, mappedAddress) {
		c.prgMemory[mappedAddress] = data
		return true
	} else {
		return false
	}
}
