package pasm

const (
	JmpConAlways  uint16 = 0b000 << 4
	JmpConNotX    uint16 = 0b001 << 4
	JmpConXDec    uint16 = 0b010 << 4
	JmpConNotY    uint16 = 0b011 << 4
	JmpConYDec    uint16 = 0b100 << 4
	JmpConXNotY   uint16 = 0b101 << 4
	JmpConPin     uint16 = 0b110 << 4
	JmpConNotOsre uint16 = 0b111 << 4
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
