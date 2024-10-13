package main

import (
	"fmt"
)

func numbers() {
	fmt.Println("Solving numbers")
	integerOperations()
	floatOperations()
	mixedOperations()
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

	cond := a.GT(b)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func integerOperations2() {
	fmt.Println("Solving integerOperations case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond := a.GT(b)
	ctx.Solver.Assert(prevCond.Not())

	cond := a.LT(b)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func integerOperations3() {
	fmt.Println("Solving integerOperations case 2")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	prevCond1 := a.GT(b)
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := a.LE(b)
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
	// todo Error occurred: smt tactic failed to show goal to be sat/unsat (incomplete (theory arithmetic))
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

	cond1 := a.Mod(two).Eq(zero)
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

	cond1 := a.Mod(two).Eq(zero)
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

	prevCond1 := a.Mod(two).Eq(zero)
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

	prevCond1 := a.Mod(two).Eq(zero)
	ctx.Solver.Assert(prevCond1.Not())

	result := ctx.float(a).Add(b)

	ten := ctx.floatConst(10.0)

	prevCond2 := result.LT(ten)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}
