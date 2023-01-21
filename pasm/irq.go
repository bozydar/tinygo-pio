package pasm

const (
	IrqModeNoWait = 0
	IrqModeBlock  = 0b01 << 5
	IrqModeClear  = 0b10 << 5
)

type Irq struct {
	Mode    uint16
	Index   uint16
	Delay   uint16
	SideSet uint16
}

const irqCode uint16 = 0b110 << 13

func (irq *Irq) ToBinary(b *Builder) uint16 {
	return irqCode | b.evalDss(irq.Delay, irq.SideSet) | irq.Mode | irq.Index
}
