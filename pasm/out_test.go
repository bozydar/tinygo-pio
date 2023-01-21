package pasm

import "testing"

func Test_Out(t *testing.T) {
	t.Run("Produces OUT instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Out{Dst: OutDstX, BitCount: 2})

		if result := b.ToHex(); result != "6022\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces OUT instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Out{Dst: OutDstIsr, BitCount: 4})

		if result := b.ToHex(); result != "60c4\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
