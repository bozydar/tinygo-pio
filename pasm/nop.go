package pasm

type Nop struct {
	Delay   uint16
	SideSet uint16
}

func (nop *Nop) ToBinary(b *Builder) uint16 {
	return 0xa042 | b.evalDss(nop.Delay, nop.SideSet)
}
