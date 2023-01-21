package pasm

import "testing"

func Test_Push(t *testing.T) {
	t.Run("Produces PUSH instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Push{})

		if result := b.ToHex(); result != "8020\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces PUSH instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Push{IfFull: true})

		if result := b.ToHex(); result != "8060\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces PUSH instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Push{IfFull: true, NoBlock: true})

		if result := b.ToHex(); result != "8040\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces PUSH instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Push{NoBlock: true})

		if result := b.ToHex(); result != "8000\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
