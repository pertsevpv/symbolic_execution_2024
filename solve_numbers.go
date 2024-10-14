package main

import (
	"fmt"
)

func numbers() {
	fmt.Println("Solving numbers")
	integerOperations()
	floatOperations()
	mixedOperations()
	nestedConditions()
	bitwiseOperations()
	advancedBitwise()
	combinedBitwise()
	nestedBitwise()
}

func integerOperations() {
	fmt.Println("Solving integerOperations")
	integerOperations1()
	integerOperations2()
	integerOperations3()
}

func integerOperations1() {
	fmt.Println("Solving integerOperations case 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	cond := a.SGT(b)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func integerOperations2() {
	fmt.Println("Solving integerOperations case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond := a.SGT(b)
	ctx.Solver.Assert(prevCond.Not())

	cond := a.SLT(b)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func integerOperations3() {
	fmt.Println("Solving integerOperations case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond1 := a.SGT(b)
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := a.SLT(b)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func floatOperations() {
	fmt.Println("Solving floatOperations")
	floatOperations1()
	floatOperations2()
	floatOperations3()
}

func floatOperations1() {
	fmt.Println("Solving floatOperations case 1")

	ctx := newContext()
	x := ctx.newFloat("x")
	y := ctx.newFloat("y")

	cond := x.GT(y)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func floatOperations2() {
	fmt.Println("Solving floatOperations case 2")

	ctx := newContext()
	x := ctx.newFloat("x")
	y := ctx.newFloat("y")

	prevCond := x.GT(y)
	ctx.Solver.Assert(prevCond.Not())

	cond := x.LT(y)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func floatOperations3() {
	fmt.Println("Solving floatOperations case 2")

	ctx := newContext()
	x := ctx.newFloat("x")
	y := ctx.newFloat("y")

	prevCond1 := x.GT(y)
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := x.LT(y)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func mixedOperations() {
	fmt.Println("Solving mixedOperations")
	mixedOperations11()
	mixedOperations12()
	mixedOperations21()
	mixedOperations22()
}

func mixedOperations11() {
	fmt.Println("Solving mixedOperations cases 1 & 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newFloat("b")
	zero := ctx.intConst(0)
	two := ctx.intConst(2)

	cond1 := a.SMod(two).Eq(zero)
	ctx.Solver.Assert(cond1)

	result := ctx.float(a).Add(b)

	ten := ctx.floatConst(10.0)

	cond2 := result.LT(ten)
	ctx.Solver.Assert(cond2)

	ctx.Check()
}

func mixedOperations12() {
	fmt.Println("Solving mixedOperations cases 1 & 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newFloat("b")
	zero := ctx.intConst(0)
	two := ctx.intConst(2)

	cond1 := a.SMod(two).Eq(zero)
	ctx.Solver.Assert(cond1)

	result := ctx.float(a).Add(b)

	ten := ctx.floatConst(10.0)

	prevCond2 := result.LT(ten)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func mixedOperations21() {
	fmt.Println("Solving mixedOperations cases 2 & 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newFloat("b")
	zero := ctx.intConst(0)
	two := ctx.intConst(2)

	prevCond1 := a.SMod(two).Eq(zero)
	ctx.Solver.Assert(prevCond1.Not())

	result := ctx.float(a).Add(b)

	ten := ctx.floatConst(10.0)

	cond2 := result.LT(ten)
	ctx.Solver.Assert(cond2)

	ctx.Check()
}

func mixedOperations22() {
	fmt.Println("Solving mixedOperations cases 2 & 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newFloat("b")
	zero := ctx.intConst(0)
	two := ctx.intConst(2)

	prevCond1 := a.SMod(two).Eq(zero)
	ctx.Solver.Assert(prevCond1.Not())

	result := ctx.float(a).Add(b)

	ten := ctx.floatConst(10.0)

	prevCond2 := result.LT(ten)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func nestedConditions() {
	fmt.Println("Solving nestedConditions")
	nestedConditions1()
	nestedConditions2()
	nestedConditions3()
}

func nestedConditions1() {
	fmt.Println("Solving nestedConditions case 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newFloat("b")

	cond1 := a.SLT(ctx.intConst(0))
	ctx.Solver.Assert(cond1)

	cond2 := b.LT(ctx.floatConst(0.0))
	ctx.Solver.Assert(cond2)

	ctx.Check()
}

func nestedConditions2() {
	fmt.Println("Solving nestedConditions case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newFloat("b")

	cond1 := a.SLT(ctx.intConst(0))
	ctx.Solver.Assert(cond1)

	prevCond2 := b.LT(ctx.floatConst(0.0))
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func nestedConditions3() {
	fmt.Println("Solving nestedConditions case 3")

	ctx := newContext()
	a := ctx.newInt("a")

	prevCond := a.SLT(ctx.intConst(0))
	ctx.Solver.Assert(prevCond.Not())

	ctx.Check()
}

func bitwiseOperations() {
	fmt.Println("Solving bitwiseOperations")
	bitwiseOperations1()
	bitwiseOperations2()
	bitwiseOperations3()
}

func bitwiseOperations1() {
	fmt.Println("Solving bitwiseOperations case 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")
	zero := ctx.intConst(0)
	one := ctx.intConst(1)

	cond := (a.And(one).Eq(zero)).And(b.And(one).Eq(zero))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func bitwiseOperations2() {
	fmt.Println("Solving bitwiseOperations case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")
	zero := ctx.intConst(0)
	one := ctx.intConst(1)

	prevCond1 := (a.And(one).Eq(zero)).And(b.And(one).Eq(zero))
	ctx.Solver.Assert(prevCond1.Not())

	cond := (a.And(one).Eq(one)).And(b.And(one).Eq(one))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func bitwiseOperations3() {
	fmt.Println("Solving bitwiseOperations case 3")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")
	zero := ctx.intConst(0)
	one := ctx.intConst(1)

	prevCond1 := (a.And(one).Eq(zero)).And(b.And(one).Eq(zero))
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := (a.And(one).Eq(one)).And(b.And(one).Eq(one))
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func advancedBitwise() {
	fmt.Println("Solving advancedBitwise")
	advancedBitwise1()
	advancedBitwise2()
	advancedBitwise3()
}

func advancedBitwise1() {
	fmt.Println("Solving advancedBitwise case 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	cond := a.SGT(b)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func advancedBitwise2() {
	fmt.Println("Solving advancedBitwise case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond := a.SGT(b)
	ctx.Solver.Assert(prevCond.Not())

	cond := a.SLT(b)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func advancedBitwise3() {
	fmt.Println("Solving advancedBitwise case 3")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond1 := a.SGT(b)
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := a.SLT(b)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func combinedBitwise() {
	fmt.Println("Solving combinedBitwise")
	combinedBitwise1()
	combinedBitwise2()
	combinedBitwise3()
}

func combinedBitwise1() {
	fmt.Println("Solving combinedBitwise case 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	cond := a.And(b).Eq(ctx.intConst(0))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func combinedBitwise2() {
	fmt.Println("Solving combinedBitwise case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond := a.And(b).Eq(ctx.intConst(0))
	ctx.Solver.Assert(prevCond.Not())

	result := a.And(b)

	cond := result.SGT(ctx.intConst(10))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func combinedBitwise3() {
	fmt.Println("Solving combinedBitwise case 3")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond1 := a.And(b).Eq(ctx.intConst(0))
	ctx.Solver.Assert(prevCond1.Not())

	result := a.And(b)

	prevCond2 := result.SGT(ctx.intConst(10))
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func nestedBitwise() {
	fmt.Println("Solving nestedBitwise")
	nestedBitwise1()
	nestedBitwise2()
	nestedBitwise3()
	nestedBitwise4()
}

func nestedBitwise1() {
	fmt.Println("Solving nestedBitwise case 1")

	ctx := newContext()
	a := ctx.newInt("a")

	cond := a.SLT(ctx.intConst(0))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func nestedBitwise2() {
	fmt.Println("Solving nestedBitwise case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")
	zero := ctx.intConst(0)

	prevCond1 := a.SLT(zero)
	ctx.Solver.Assert(prevCond1)

	cond := b.SLT(zero)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func nestedBitwise3() {
	fmt.Println("Solving nestedBitwise case 3")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")
	zero := ctx.intConst(0)

	prevCond1 := a.SLT(zero)
	ctx.Solver.Assert(prevCond1)

	prevCond2 := b.SLT(zero)
	ctx.Solver.Assert(prevCond2)

	cond := a.And(b).Eq(zero)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func nestedBitwise4() {
	fmt.Println("Solving nestedBitwise case 4")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")
	zero := ctx.intConst(0)

	prevCond1 := a.SLT(zero)
	ctx.Solver.Assert(prevCond1)

	prevCond2 := b.SLT(zero)
	ctx.Solver.Assert(prevCond2)

	prevCond3 := a.And(b).Eq(zero)
	ctx.Solver.Assert(prevCond3.Not())

	ctx.Check()
}
