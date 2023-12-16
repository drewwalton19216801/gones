package cpu

// Bus is an interface for interacting with the CPU's memory
type Bus interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}
