package pasm

import (
	"testing"
)

func TestBuilder_Label(t *testing.T) {
	t.Run("Adds label.", func(t *testing.T) {
		b := NewBuilder()
		b.Label("entry_point")
		b.Add(&Pull{})
		b.Add(&Set{Dst: SetDstX, Data: 23})
		b.Label("bitloop")
		b.Add(&Set{Dst: SetDstPins, Data: 1})
		b.Add(&Out{Dst: OutDstY, BitCount: 1, Delay: 5})
		b.Add(&Jmp{Con: JmpConNotY, Label: "skip"})
		b.Add(&Nop{Delay: 5})

		if b.Labels["loop"] != 0 {
			t.Error("Label not added")
		}
	})
}
