package stack

import (
	"math/big"
	"testing"

	"github.com/xssnick/tonutils-go/tvm/vm"
)

func TestOVER2Operation(t *testing.T) {
	st := vm.NewStack()

	st.PushAny(big.NewInt(1))
	st.PushAny(big.NewInt(2))
	st.PushAny(big.NewInt(3))
	st.PushAny(big.NewInt(4))
	over2 := OVER2()

	err := over2.Interpret(&vm.State{
		Stack: st,
	})
	if err != nil {
		t.Fatal("Failed ROT interpretation:", err.Error())
	}

	b, err := st.PopInt()
	if err != nil || b.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("Expected 2 at b, got %v", b)
	}
	a, err := st.PopInt()
	if err != nil || a.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("Expected 1 at a, got %v", a)
	}
	d, err := st.PopInt()
	if err != nil || d.Cmp(big.NewInt(4)) != 0 {
		t.Errorf("Expected 4 at d, got %v", d)
	}
	c, err := st.PopInt()
	if err != nil || c.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("Expected 3 at c, got %v", c)
	}
	b, err = st.PopInt()
	if err != nil || b.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("Expected 2 at b, got %v", b)
	}
	a, err = st.PopInt()
	if err != nil || a.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("Expected 1 at a, got %v", a)
	}
}
