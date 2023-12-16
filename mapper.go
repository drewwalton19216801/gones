package main

type Mapper interface {
	// Transform CPU bus address to PRG ROM address
	cpuMapRead(addr uint16, mappedAddress *uint32) bool
	cpuMapWrite(addr uint16, mappedAddress *uint32) bool
	// Transform PPU bus address to CHR ROM offset
	ppuMapRead(addr uint16, mappedAddress *uint32) bool
	ppuMapWrite(addr uint16, mappedAddress *uint32) bool
}
