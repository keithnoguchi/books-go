package ch04

type UUIDCounter map[string]int

func NewUUIDCounter() *UUIDCounter {
	return &UUIDCounter{}
}

func (c *UUIDCounter) Count(id []byte) {
	(*c)[string(id)]++
}
