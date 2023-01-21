package pasm

const (
	SetDstPins    uint16 = 0b000 << 5
	SetDstX       uint16 = 0b001 << 5
	SetDstY       uint16 = 0b010 << 5
	SetDstPinDirs uint16 = 0b100 << 5
)

type Set struct {
	Dst     uint16
	Data    uint16
	Delay   uint16
	SideSet uint16
}

const setCode uint16 = 0b111 << 13

func (set *Set) ToBinary(b *Builder) uint16 {
	return setCode | b.evalDss(set.Delay, set.SideSet) | set.Dst | set.Data
}
