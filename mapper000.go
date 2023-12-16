package main

type Mapper000 struct {
	prgBanks uint8
	chrBanks uint8
}

func (m *Mapper000) cpuMapRead(addr uint16, mappedAddress *uint32) bool {
	if addr >= 0x8000 && addr <= 0xFFFF {
		if uint16(m.prgBanks) > 1 {
			*mappedAddress = uint32(addr & 0x7FFF)
		} else {
			*mappedAddress = uint32(addr & 0x3FFF)
		}
		return true
	}
	return false
}

func (m *Mapper000) cpuMapWrite(addr uint16, mappedAddress *uint32) bool {
	if addr >= 0x8000 && addr <= 0xFFFF {
		if m.prgBanks > 1 {
			*mappedAddress = uint32(addr & 0x7FFF)
		} else {
			*mappedAddress = uint32(addr & 0x3FFF)
		}
		return true
	}
	return false
}

func (m *Mapper000) ppuMapRead(addr uint16, mappedAddress *uint32) bool {
	return true
}

func (m *Mapper000) ppuMapWrite(addr uint16, mappedAddress *uint32) bool {
	return true
}
