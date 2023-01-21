package pasm

import "testing"

func Test_Set(t *testing.T) {
	t.Run("Produces SET instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Set{Dst: SetDstPins, Data: 0b111})

		if result := b.ToHex(); result != "e007\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces SET instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Set{Dst: SetDstPinDirs, Data: 0b111})

		if result := b.ToHex(); result != "e087\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces SET instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Set{Dst: SetDstX, Data: 0b111})

		if result := b.ToHex(); result != "e027\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces SET instruction. Case 4.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Set{Dst: SetDstY, Data: 0b111})

		if result := b.ToHex(); result != "e047\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
