package pasm

import "testing"

func Test_Pull(t *testing.T) {
	t.Run("Produces PULL instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Label("loop")
		b.Add(&Pull{})

		if result := b.ToHex(); result != "80a0\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces PULL instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Label("loop")
		b.Add(&Pull{IfEmpty: true, NoBlock: true})

		if result := b.ToHex(); result != "80c0\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces PULL instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder()
		b.Label("loop")
		b.Add(&Pull{IfEmpty: true, NoBlock: false})

		if result := b.ToHex(); result != "80e0\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces PULL instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder()
		b.Label("loop")
		b.Add(&Pull{IfEmpty: false, NoBlock: true})

		if result := b.ToHex(); result != "8080\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
