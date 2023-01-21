package pasm

const (
	ifFull uint16 = 0b10 << 5
)

type Push struct {
	IfFull  bool
	NoBlock bool
	Delay   uint16
	SideSet uint16
}

const pushCode = pullCode

func (pull *Push) ToBinary(b *Builder) uint16 {
	result := pushCode | b.evalDss(pull.Delay, pull.SideSet)
	if pull.IfFull {
		result |= ifFull
	}
	if !pull.NoBlock {
		result |= block
	}
	return result
}
