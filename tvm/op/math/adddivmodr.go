package math

import (
	"github.com/xssnick/tonutils-go/tvm/op/helpers"
	"github.com/xssnick/tonutils-go/tvm/vm"
	"github.com/xssnick/tonutils-go/tvm/vmerr"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return ADDDIVMODR() })
}

func ADDDIVMODR() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			z, err := state.Stack.PopIntFinite()
			if err != nil {
				return err
			}
			w, err := state.Stack.PopIntFinite()
			if err != nil {
				return err
			}
			x, err := state.Stack.PopIntFinite()
			if err != nil {
				return err
			}

			if z.Sign() == 0 {
				// division by 0
				return vmerr.VMError{
					Code: vmerr.ErrIntOverflow.Code,
					Msg:  "division by zero",
				}
			}

			sum := w.Add(x, w)
			q := helpers.DivRound(sum, z)
			r := x.Sub(sum, z.Mul(z, q))

			err = state.Stack.PushInt(q)
			if err != nil {
				return err
			}

			return state.Stack.PushInt(r)
		},
		Name:   "ADDDIVMODR",
		Prefix: []byte{0xA9, 0x01},
	}
}
