package runtime

import (
	"github.com/Zac-Garby/radon/bytecode"
	"github.com/Zac-Garby/radon/object"
)

// A Frame is created for each function call, and also one for the main program. It contains
// the bytecode to execute, along with the frame's constants and names, and other data.
type Frame struct {
	prev          *Frame
	code          bytecode.Code
	offset        int
	vm            *VM
	store         *Store
	stack         *Stack
	breaks, nexts []int
	constants     []object.Object
	names         []string
}

func (f *Frame) offsetToInstructionIndex(offset int) int {
	var index, counter int

	for _, instr := range f.code {
		if bytecode.Instructions[instr.Code].HasArg {
			counter += 3
		} else {
			counter++
		}

		if counter >= offset {
			return index
		}

		index++
	}

	return index
}

func (f *Frame) getName(arg rune) (string, bool) {
	index := int(arg)
	if index < len(f.names) {
		return f.names[index], true
	}

	return "", false
}

func (f *Frame) searchName(name string) (object.Object, bool) {
	if val, ok := f.store.Get(name); ok {
		return val, true
	}

	return nil, false
}