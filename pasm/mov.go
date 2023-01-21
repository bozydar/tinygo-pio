package pasm

const (
	MovOpNone    uint16 = 0b00
	MovOpInvert  uint16 = 0b01 << 3
	MovOpReverse uint16 = 0b10 << 3

	MovDstPins uint16 = 0b000
	MovDstX    uint16 = 0b001 << 5
	MovDstY    uint16 = 0b010 << 5
	MovDstExec uint16 = 0b100 << 5
	MovDstPc   uint16 = 0b101 << 5
	MovDstIsr  uint16 = 0b110 << 5
	MovDstOsr  uint16 = 0b111 << 5

	MovSrcPins   uint16 = 0b000
	MovSrcX      uint16 = 0b001
	MovSrcY      uint16 = 0b010
	MovSrcNull   uint16 = 0b011
	MovSrcStatus uint16 = 0b101
	MovSrcIsr    uint16 = 0b110
	MovSrcOsr    uint16 = 0b111
)

type Mov struct {
	Dst     uint16
	Src     uint16
	Op      uint16
	Delay   uint16
	SideSet uint16
}

const movCode uint16 = 0b101 << 13

func (mov *Mov) ToBinary(b *Builder) uint16 {
	return movCode | b.evalDss(mov.Delay, mov.SideSet) | mov.Dst | mov.Op | mov.Src
}
