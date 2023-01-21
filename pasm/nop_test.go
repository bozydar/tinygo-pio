package pasm

import "testing"

func Test_Nop(t *testing.T) {
	t.Run("Produces NOP instruction. Case 1.", func(t *testing.T) {
		b := NewBuilder()
		b.Add(&Nop{Delay: 3})

		if result := b.ToHex(); result != "a342\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces NOP instruction. Case 2.", func(t *testing.T) {
		b := NewBuilder().SideSet(2, false)
		b.Add(&Nop{Delay: 3})

		if result := b.ToHex(); result != "a342\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces NOP instruction. Case 3.", func(t *testing.T) {
		b := NewBuilder().SideSet(3, false)
		b.Add(&Nop{SideSet: 2, Delay: 1})

		if result := b.ToHex(); result != "a942\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces NOP instruction. Case 4.", func(t *testing.T) {
		b := NewBuilder().SideSet(3, false)
		b.Add(&Nop{SideSet: 3, Delay: 2})

		if result := b.ToHex(); result != "ae42\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces NOP instruction. Case 5.", func(t *testing.T) {
		b := NewBuilder().SideSet(2, true)
		b.Add(&Nop{SideSet: 3, Delay: 2})

		// 1011 1110 0100 0010 - be42
		t.Logf("%08b", 0xbe42)
		t.Logf("%08b", b.ToBinary()[0])
		if result := b.ToHex(); result != "be42\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})

	t.Run("Produces NOP instruction. Case 6.", func(t *testing.T) {
		b := NewBuilder().SideSet(1, true)
		b.Add(&Nop{SideSet: 1, Delay: 4})

		t.Logf("%08b", 0xbc42)
		t.Logf("%08b", b.ToBinary()[0])
		if result := b.ToHex(); result != "bc42\n" {
			t.Errorf("Hex doesn't match `%s`", result)
		}
	})
}
