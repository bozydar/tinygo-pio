package pasm

const (
	OutDstPins    uint16 = 0b000
	OutDstX       uint16 = 0b001 << 5
	OutDstY       uint16 = 0b010 << 5
	OutDstNull    uint16 = 0b011 << 5
	OutDstPinDirs uint16 = 0b100 << 5
	OutDstPc      uint16 = 0b101 << 5
	OutDstIsr     uint16 = 0b110 << 5
	OutDstExec    uint16 = 0b111 << 5
)

type Out struct {
	Dst      uint16
	BitCount uint16
	Delay    uint16
	SideSet  uint16
}

const outCode uint16 = 0b011 << 13

func (in *Out) ToBinary(b *Builder) uint16 {
	return outCode | b.evalDss(in.Delay, in.SideSet) | in.Dst | in.BitCount
}
