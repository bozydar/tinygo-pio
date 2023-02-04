package pasm

const (
	JmpConAlways  uint16 = 0b000 << 5
	JmpConNotX    uint16 = 0b001 << 5
	JmpConXDec    uint16 = 0b010 << 5
	JmpConNotY    uint16 = 0b011 << 5
	JmpConYDec    uint16 = 0b100 << 5
	JmpConXNotY   uint16 = 0b101 << 5
	JmpConPin     uint16 = 0b110 << 5
	JmpConNotOsre uint16 = 0b111 << 5
)

type Jmp struct {
	Con     uint16
	Label   string
	Delay   uint16
	SideSet uint16
}

func (jmp *Jmp) ToBinary(b *Builder) uint16 {
	addr := b.Labels[jmp.Label]
	return b.evalDss(jmp.Delay, jmp.SideSet) |
		jmp.Con | addr
}
