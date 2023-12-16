package cpu

// Bus is an interface for interacting with the memory attached to the CPU
type Bus interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
}
