package pasm

import (
	"fmt"
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

	t.Run("Adds label.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Pull{NoBlock: false})
		b.Add(&Out{Dst: OutDstY, BitCount: 32})
		b.Label("wrap_target")
		b.Add(&Mov{Dst: MovDstX, Src: MovSrcY})
		b.Add(&Set{Dst: SetDstPins, Data: 1})
		b.Label("dec_one")
		b.Add(&Jmp{Con: JmpConXDec, Label: "dec_one"})
		b.Add(&Mov{Dst: MovDstX, Src: MovSrcY})
		b.Add(&Set{Dst: SetDstPins, Data: 0})
		b.Label("dec_two")
		b.Add(&Jmp{Con: JmpConXDec, Label: "dec_two"})

		fmt.Println(b.ToHex())
	})
}
