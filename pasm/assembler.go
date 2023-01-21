package pasm

import (
	"fmt"
)

type Builder struct {
	Labels       map[string]uint16
	Instructions []Instruction
	SideSetCount uint16
	SideSetOpt   bool
	BaseAddr     uint16
}

type Instruction interface {
	ToBinary(b *Builder) uint16
}

func (b *Builder) evalDss(delay, sideSet uint16) uint16 {
	siteSetCount := b.SideSetCount
	if b.SideSetOpt {
		sideSet = (sideSet << 1) | 1
		siteSetCount++
	}
	return ((sideSet << (5 - siteSetCount)) | delay) << 8
}

func NewBuilder() *Builder {
	// TODO inject variable names
	// TODO curAddr can be different at the very beginning
	// look at ".origin <offset>" p.317
	return &Builder{
		BaseAddr:     0,
		SideSetCount: 0,
		Instructions: make([]Instruction, 0),
		Labels:       make(map[string]uint16),
	}
}

func (b *Builder) SideSet(count uint16, opt bool) *Builder {
	b.SideSetCount = count
	b.SideSetOpt = opt
	return b
}

func (b *Builder) Label(name string) {
	b.Labels[name] = b.curAddr()
}

func (b *Builder) Add(i Instruction) {
	b.Instructions = append(b.Instructions, i)
}

func (b *Builder) curAddr() uint16 {
	return uint16(len(b.Instructions))
}

func (b *Builder) ToBinary() (result []uint16) {
	result = make([]uint16, 0)
	for _, i := range b.Instructions {
		result = append(result, i.ToBinary(b))
	}

	return
}

func (b *Builder) ToHex() (result string) {
	for _, v := range b.ToBinary() {
		result += fmt.Sprintf("%04x\n", v)
	}
	return
}
