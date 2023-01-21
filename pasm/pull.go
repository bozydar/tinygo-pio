package pasm

const (
	ifEmpty uint16 = 0b10 << 5
	block   uint16 = 0b01 << 5
)

type Pull struct {
	IfEmpty bool
	NoBlock bool
	Delay   uint16
	SideSet uint16
}

const pullCode uint16 = 0b100 << 13

func (pull *Pull) ToBinary(b *Builder) uint16 {
	result := pullCode | b.evalDss(pull.Delay, pull.SideSet) | (1 << 7)
	if pull.IfEmpty {
		result |= ifEmpty
	}
	if !pull.NoBlock {
		result |= block
	}
	return result
}
