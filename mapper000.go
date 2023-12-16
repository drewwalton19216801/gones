package main

type Mapper000 struct {
}

func (m *Mapper000) cpuMapRead(addr uint16, mappedAddress uint32) bool {
	return true
}

func (m *Mapper000) cpuMapWrite(addr uint16, mappedAddress uint32) bool {
	return true
}

func (m *Mapper000) ppuMapRead(addr uint16, mappedAddress uint32) bool {
	return true
}

func (m *Mapper000) ppuMapWrite(addr uint16, mappedAddress uint32) bool {
	return true
}
