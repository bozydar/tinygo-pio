package pasm

import "testing"

func Test_Jmp(t *testing.T) {
	t.Run("Produces JMP instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Label("loop")
		b.Add(&Jmp{Label: "loop"})

		if result := b.ToHex(); result != "0000\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces JMP instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Jmp{Label: "loop"})
		b.Add(&Mov{Src: MovSrcY, Dst: MovDstY})
		b.Label("loop")
		b.Add(&Mov{Src: MovSrcY, Dst: MovDstY})

		want := `0002
a042
a042
`
		if result := b.ToHex(); result != want {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
