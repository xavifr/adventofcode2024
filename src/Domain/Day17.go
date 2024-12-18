package Domain

import (
	"math"
)

type D17State string

const (
	D17_RUNNING D17State = "RUNNING"
	D17_HALTED  D17State = "HALTED"
	D17_PANIC   D17State = "PANIC"
)

type D17Operation int

const (
	D17_ADV D17Operation = 0
	D17_BXL D17Operation = 1
	D17_BST D17Operation = 2
	D17_JNZ D17Operation = 3
	D17_BXC D17Operation = 4
	D17_OUT D17Operation = 5
	D17_BDV D17Operation = 6
	D17_CDV D17Operation = 7
)

type D17OperationFunc func(int, int) int

type D17Machine struct {
	RegA, RegB, RegC int64
	IP               int
	Program          []int
	State            D17State
	Output           []int
}

func NewD17Machine() D17Machine {
	return D17Machine{
		Program: make([]int, 0),
		Output:  make([]int, 0),
		State:   D17_RUNNING,
	}
}

func (d *D17Machine) Step() (int, D17State) {
	inst, oper, state := d.nextOp()
	if state != D17_RUNNING {
		return d.IP, d.State
	}

	switch inst {
	case D17_ADV:
		d.opADV(oper)
	case D17_BXL:
		d.opBXL(oper)
	case D17_BST:
		d.opBST(oper)
	case D17_JNZ:
		d.opJNZ(oper)
	case D17_BXC:
		d.opBXC(oper)
	case D17_OUT:
		d.opOUT(oper)
	case D17_BDV:
		d.opBDV(oper)
	case D17_CDV:
		d.opCDV(oper)
	}

	return d.IP, d.State
}

func (d *D17Machine) Run() (int, D17State) {
	for {
		ip, state := d.Step()
		if state != D17_RUNNING {
			return ip, state
		}
	}
}

func (d *D17Machine) GetOutput() []int {
	return d.Output
}

func (d *D17Machine) opADV(oper int) {
	d.incInstructionRegister()
	num := d.RegA
	comboOper, state := d.getComboOperand(oper)
	if state != D17_RUNNING {
		return
	}

	den := math.Pow(float64(2), float64(comboOper))

	d.RegA = int64(num / int64(den))
}

func (d *D17Machine) opBXL(oper int) {
	d.incInstructionRegister()
	opA := d.RegB
	opB, state := d.getLiteralOperand(oper)
	if state != D17_RUNNING {
		return
	}

	d.RegB = opA ^ int64(opB)
}

func (d *D17Machine) opBST(oper int) {
	d.incInstructionRegister()
	opA, state := d.getComboOperand(oper)
	if state != D17_RUNNING {
		return
	}

	d.RegB = int64(opA % 8)
}
func (d *D17Machine) opJNZ(oper int) {
	if d.RegA == 0 {
		d.incInstructionRegister()
		return
	}

	opA, state := d.getLiteralOperand(oper)
	if state != D17_RUNNING {
		d.incInstructionRegister()
		return
	}

	d.IP = opA
}
func (d *D17Machine) opBXC(_ int) {
	d.incInstructionRegister()
	opA := d.RegB
	opB := d.RegC

	d.RegB = opA ^ opB
}

func (d *D17Machine) opOUT(oper int) {
	d.incInstructionRegister()
	opA, state := d.getComboOperand(oper)
	if state != D17_RUNNING {
		return
	}

	d.Output = append(d.Output, opA%8)
}
func (d *D17Machine) opBDV(oper int) {
	d.incInstructionRegister()
	num := d.RegA
	comboOper, state := d.getComboOperand(oper)
	if state != D17_RUNNING {
		return
	}

	den := math.Pow(float64(2), float64(comboOper))

	d.RegB = int64(num / int64(den))
}
func (d *D17Machine) opCDV(oper int) {
	d.incInstructionRegister()
	num := d.RegA
	comboOper, state := d.getComboOperand(oper)
	if state != D17_RUNNING {
		return
	}

	den := math.Pow(float64(2), float64(comboOper))

	d.RegC = int64(num / int64(den))
}

func (d *D17Machine) getComboOperand(operand int) (int, D17State) {
	if operand >= 0 && operand <= 3 {
		return operand, d.State
	}

	switch operand {
	case 4:
		return int(d.RegA % 8), d.State
	case 5:
		return int(d.RegB % 8), d.State
	case 6:
		return int(d.RegC % 8), d.State
	case 7:
		d.State = D17_PANIC
		return 0, d.State
	default:
		d.State = D17_PANIC
		return 0, d.State
	}
}
func (d *D17Machine) getLiteralOperand(operand int) (int, D17State) {
	return operand % 8, d.State
}

func (d *D17Machine) incInstructionRegister() D17State {
	d.IP += 2

	return d.State
}

func (d *D17Machine) nextOp() (D17Operation, int, D17State) {
	if len(d.Program) < d.IP+2 {
		d.State = D17_HALTED
	}

	if d.State != D17_RUNNING {
		return 0, 0, d.State
	}

	inst := d.Program[d.IP]
	oper := d.Program[d.IP+1]

	return D17Operation(inst), oper, d.State
}
