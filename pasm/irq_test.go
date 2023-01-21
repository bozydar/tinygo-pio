package pasm

import "testing"

func Test_Irq(t *testing.T) {
	t.Run("Produces IRQ instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Irq{Mode: IrqModeBlock, Index: 1})

		if result := b.ToHex(); result != "c021\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces IRQ instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Irq{Mode: IrqModeNoWait, Index: 1})

		if result := b.ToHex(); result != "c001\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces IRQ instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Irq{Mode: IrqModeClear, Index: 1})

		if result := b.ToHex(); result != "c041\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
