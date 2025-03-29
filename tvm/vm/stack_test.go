package vm

import (
	"math/big"
	"testing"
)

func Test_StackRotate(t *testing.T) {
	t.Run("123 -> 231", func(t *testing.T) {
		stack := NewStack()

		stack.PushAny(big.NewInt(1))
		stack.PushAny(big.NewInt(2))
		stack.PushAny(big.NewInt(3))

		xy, err := stack.FromTop(3)
		if err != nil {
			t.Error(err)
		}
		y, err := stack.FromTop(2)
		if err != nil {
			t.Error(err)
		}
		if err := stack.Rotate(xy, y, stack.Len()); err != nil {
			t.Error(err)
		}

		a, err := stack.PopInt()
		if err != nil || a.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Expected 1 at a, got %v", a)
		}
		c, err := stack.PopInt()
		if err != nil || c.Cmp(big.NewInt(3)) != 0 {
			t.Errorf("Expected 3 at c, got %v", c)
		}
		b, err := stack.PopInt()
		if err != nil || b.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("Expected 2 at b, got %v", b)
		}
	})

	t.Run("123 -> 312", func(t *testing.T) {
		stack := NewStack()

		stack.PushAny(big.NewInt(1))
		stack.PushAny(big.NewInt(2))
		stack.PushAny(big.NewInt(3))

		xy, err := stack.FromTop(3)
		if err != nil {
			t.Error(err)
		}
		y, err := stack.FromTop(1)
		if err != nil {
			t.Error(err)
		}
		if err := stack.Rotate(xy, y, stack.Len()); err != nil {
			t.Error(err)
		}

		b, err := stack.PopInt()
		if err != nil || b.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("Expected 2 at b, got %v", b)
		}
		a, err := stack.PopInt()
		if err != nil || a.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Expected 1 at a, got %v", a)
		}
		c, err := stack.PopInt()
		if err != nil || c.Cmp(big.NewInt(3)) != 0 {
			t.Errorf("Expected 3 at c, got %v", c)
		}

	})

	t.Run("1234 - 3421", func(t *testing.T) {
		stack := NewStack()

		stack.PushAny(big.NewInt(1))
		stack.PushAny(big.NewInt(2))
		stack.PushAny(big.NewInt(3))
		stack.PushAny(big.NewInt(4))

		xy, err := stack.FromTop(4)
		if err != nil {
			t.Error(err)
		}
		y, err := stack.FromTop(2)
		if err != nil {
			t.Error(err)
		}
		if err := stack.Rotate(xy, y, stack.Len()); err != nil {
			t.Error(err)
		}

		b, err := stack.PopInt()
		if err != nil || b.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("Expected 2 at b, got %v", b)
		}
		a, err := stack.PopInt()
		if err != nil || a.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Expected 1 at a, got %v", a)
		}
		d, err := stack.PopInt()
		if err != nil || d.Cmp(big.NewInt(4)) != 0 {
			t.Errorf("Expected 4 at d, got %v", d)
		}
		c, err := stack.PopInt()
		if err != nil || c.Cmp(big.NewInt(3)) != 0 {
			t.Errorf("Expected 3 at c, got %v", c)
		}
	})
}

func Test_Reverse(t *testing.T) {
	t.Run("123 -> 321", func(t *testing.T) {
		stack := NewStack()

		stack.PushAny(big.NewInt(1))
		stack.PushAny(big.NewInt(2))
		stack.PushAny(big.NewInt(3))
		if err := stack.Reverse(2, 0); err != nil {
			t.Error(err)
		}

		a, err := stack.PopInt()
		if err != nil || a.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Expected 1 at a, got %v", a)
		}
		b, err := stack.PopInt()
		if err != nil || b.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("Expected 2 at b, got %v", b)
		}
		c, err := stack.PopInt()
		if err != nil || c.Cmp(big.NewInt(3)) != 0 {
			t.Errorf("Expected 3 at c, got %v", c)
		}
	})

	t.Run("1234 -> 4321", func(t *testing.T) {
		stack := NewStack()

		stack.PushAny(big.NewInt(1))
		stack.PushAny(big.NewInt(2))
		stack.PushAny(big.NewInt(3))
		stack.PushAny(big.NewInt(4))

		if err := stack.Reverse(3, 0); err != nil {
			t.Error(err)
		}

		a, err := stack.PopInt()
		if err != nil || a.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Expected 1 at a, got %v", a)
		}
		b, err := stack.PopInt()
		if err != nil || b.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("Expected 2 at b, got %v", b)
		}
		c, err := stack.PopInt()
		if err != nil || c.Cmp(big.NewInt(3)) != 0 {
			t.Errorf("Expected 3 at c, got %v", c)
		}
		d, err := stack.PopInt()
		if err != nil || d.Cmp(big.NewInt(4)) != 0 {
			t.Errorf("Expected 4 at d, got %v", d)
		}
	})

	t.Run("1234 -> 4231", func(t *testing.T) {
		stack := NewStack()

		stack.PushAny(big.NewInt(1))
		stack.PushAny(big.NewInt(2))
		stack.PushAny(big.NewInt(3))
		stack.PushAny(big.NewInt(4))

		if err := stack.Reverse(2, 1); err != nil {
			t.Error(err)
		}

		d, err := stack.PopInt()
		if err != nil || d.Cmp(big.NewInt(4)) != 0 {
			t.Errorf("Expected 4 at d, got %v", d)
		}
		b, err := stack.PopInt()
		if err != nil || b.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("Expected 2 at b, got %v", b)
		}
		c, err := stack.PopInt()
		if err != nil || c.Cmp(big.NewInt(3)) != 0 {
			t.Errorf("Expected 3 at c, got %v", c)
		}
		a, err := stack.PopInt()
		if err != nil || a.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Expected 1 at a, got %v", a)
		}
	})
}

func TestStack_PushInt(t *testing.T) {
	stk := NewStack()
	i := new(big.Int).Lsh(big.NewInt(1), uint(256))
	i.Add(i, big.NewInt(1))
	println(i.String())
	println(i.Neg(i).String())

	// println(cell.BeginCell().MustStoreBigInt(i, 257).EndCell().String())
	println(i.Neg(i).String())
	err := stk.PushInt(i.Neg(i))
	if err == nil {
		t.Fatal("should be err")
	}
}
