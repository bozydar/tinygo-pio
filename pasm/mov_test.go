package pasm

import "testing"

func Test_Mov(t *testing.T) {
	t.Run("Produces MOVE instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Mov{Dst: MovDstX, Op: MovOpNone, Src: MovSrcPins})

		if result := b.ToHex(); result != "a020\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces MOVE instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Mov{Dst: MovDstY, Op: MovOpNone, Src: MovSrcPins})

		if result := b.ToHex(); result != "a040\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces MOVE instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Mov{Dst: MovDstY, Op: MovOpInvert, Src: MovSrcPins})

		if result := b.ToHex(); result != "a048\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces MOVE instruction. Case 4.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Mov{Dst: MovDstY, Op: MovOpReverse, Src: MovSrcPins})

		if result := b.ToHex(); result != "a050\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces MOVE instruction. Case 5.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Mov{Dst: MovDstY, Op: MovOpReverse, Src: MovSrcX})

		if result := b.ToHex(); result != "a051\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
