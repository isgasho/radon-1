package compiler

import (
	"fmt"

	"github.com/Zac-Garby/radon/ast"
	"github.com/Zac-Garby/radon/bytecode"
	"github.com/Zac-Garby/radon/object"
)

const maxRune = 1 << 16

func runeToBytes(r rune) (byte, byte) {
	var (
		low  = byte(r & 0xff)
		high = byte((r >> 8) & 0xff)
	)

	return low, high
}

func bytesToRune(low, high byte) rune {
	return (rune(high) << 8) | rune(low)
}

func (c *Compiler) addConst(val object.Object) (rune, error) {
	for i, cst := range c.Constants {
		if val.Equals(cst) {
			return rune(i), nil
		}
	}

	c.Constants = append(c.Constants, val)
	index := len(c.Constants) - 1

	if index >= maxRune {
		return 0, fmt.Errorf("compiler: you've somehow managed to use 65,536 constants, good job")
	}

	return rune(index), nil
}

func (c *Compiler) loadConst(index rune) {
	low, high := runeToBytes(index)
	c.push(bytecode.LoadConst, high, low)
}

func (c *Compiler) addAndLoad(obj object.Object) (rune, error) {
	index, err := c.addConst(obj)
	if err != nil {
		return 0, err
	}

	c.loadConst(index)

	return index, nil
}

func (c *Compiler) loadName(index rune) {
	low, high := runeToBytes(index)
	c.push(bytecode.LoadName, high, low)
}

func (c *Compiler) addName(name string) (rune, error) {
	for i, n := range c.Names {
		if name == n {
			return rune(i), nil
		}
	}

	c.Names = append(c.Names, name)
	index := len(c.Names) - 1

	return rune(index), nil
}

func (c *Compiler) compileName(name string) error {
	index, err := c.addName(name)
	if err != nil {
		return err
	}

	c.loadName(rune(index))

	return nil
}

func (c *Compiler) push(bytes ...byte) {
	c.Bytes = append(c.Bytes, bytes...)
}

func (c *Compiler) expandTuple(e *ast.Infix) []ast.Expression {
	if e.Operator == "," {
		li := []ast.Expression{e.Left}

		if rInf, ok := e.Right.(*ast.Infix); ok && rInf.Operator == "," {
			return append(li, c.expandTuple(rInf)...)
		} else {
			return append(li, e.Right)
		}
	} else {
		panic("compiler: non-tuple expression passed to expandTuple!")
	}
}

func (c *Compiler) addJump(target int) (rune, error) {
	for i, jmp := range c.Jumps {
		if jmp == target {
			return rune(i), nil
		}
	}

	c.Jumps = append(c.Jumps, target)
	index := len(c.Jumps) - 1

	if index >= maxRune {
		return 0, fmt.Errorf("compiler: you've somehow managed to use 65,536 jump targets, good job")
	}

	return rune(index), nil
}

func (c *Compiler) setJumpArg(start, target int) {
	index, err := c.addJump(target)
	if err != nil {
		panic(err)
	}

	low, high := runeToBytes(rune(index))
	c.Bytes[start+1] = high
	c.Bytes[start+2] = low
}
