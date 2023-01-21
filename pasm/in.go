package pasm

const (
	InSrcPins uint16 = 0b000
	InSrcX    uint16 = 0b001 << 5
	InSrcY    uint16 = 0b010 << 5
	InSrcNull uint16 = 0b011 << 5
	InSrcIsr  uint16 = 0b110 << 5
	InSrcOsr  uint16 = 0b111 << 5
)

type In struct {
	Src      uint16
	BitCount uint16
	Delay    uint16
	SideSet  uint16
}

const inCode uint16 = 0b010 << 13

func (in *In) ToBinary(b *Builder) uint16 {
	return inCode | b.evalDss(in.Delay, in.SideSet) | in.Src | in.BitCount
}
