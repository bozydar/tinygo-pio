package pasm

import "testing"

func Test_In(t *testing.T) {
	t.Run("Produces IN instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&In{Src: InSrcX, BitCount: 2})

		if result := b.ToHex(); result != "4022\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces IN instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&In{Src: InSrcIsr, BitCount: 4})

		if result := b.ToHex(); result != "40c4\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
