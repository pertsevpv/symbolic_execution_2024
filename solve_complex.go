package main

import "fmt"

func complexx() {
	fmt.Println("Solving complex")
	basicComplexOperations()
	complexComparison()
	complexOperations()
	nestedComplexOperations()
}

func basicComplexOperations() {
	fmt.Println("Solving basicComplexOperations")
	basicComplexOperations1()
	basicComplexOperations2()
	basicComplexOperations3()
}

func basicComplexOperations1() {
	fmt.Println("Solving basicComplexOperations case 1")

	ctx := newContext()
	ar, _ := ctx.newComplex("a")
	br, _ := ctx.newComplex("b")

	cond := ar.GT(br)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func basicComplexOperations2() {
	fmt.Println("Solving basicComplexOperations case 2")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")

	prevCond1 := ar.GT(br)
	ctx.Solver.Assert(prevCond1.Not())

	cond := ai.GT(bi)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func basicComplexOperations3() {
	fmt.Println("Solving basicComplexOperations case 3")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")

	prevCond1 := ar.GT(br)
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := ai.GT(bi)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func complexComparison() {
	fmt.Println("Solving complexComparison")
	complexComparison1()
	complexComparison2()
	complexComparison3()
}

func complexComparison1() {
	fmt.Println("Solving complexComparison case 1")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")

	magA := ar.Mul(ar).Add(ai.Mul(ai))
	magB := br.Mul(br).Add(bi.Mul(bi))

	cond := magA.GT(magB)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func complexComparison2() {
	fmt.Println("Solving complexComparison case 2")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")

	magA := ar.Mul(ar).Add(ai.Mul(ai))
	magB := br.Mul(br).Add(bi.Mul(bi))

	prevCond := magA.GT(magB)
	ctx.Solver.Assert(prevCond.Not())

	cond := magA.LT(magB)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func complexComparison3() {
	fmt.Println("Solving complexComparison case 3")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")

	magA := ar.Mul(ar).Add(ai.Mul(ai))
	magB := br.Mul(br).Add(bi.Mul(bi))

	prevCond1 := magA.GT(magB)
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := magA.LT(magB)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}

func complexOperations() {
	fmt.Println("Solving complexOperations")
	complexOperations1()
	complexOperations2()
	complexOperations3()
	complexOperations4()
}

func complexOperations1() {
	fmt.Println("Solving complexOperations case 1")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	zero := ctx.floatConst(0.0)

	cond := ar.Eq(zero).And(ai.Eq(zero))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func complexOperations2() {
	fmt.Println("Solving complexOperations case 2")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")
	zero := ctx.floatConst(0.0)

	prevCond := ar.Eq(zero).And(ai.Eq(zero))
	ctx.Solver.Assert(prevCond.Not())

	cond := br.Eq(zero).And(bi.Eq(zero))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func complexOperations3() {
	fmt.Println("Solving complexOperations case 3")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")
	zero := ctx.floatConst(0.0)

	prevCond1 := ar.Eq(zero).And(ai.Eq(zero))
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := br.Eq(zero).And(bi.Eq(zero))
	ctx.Solver.Assert(prevCond2.Not())

	cond := ar.GT(br)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func complexOperations4() {
	fmt.Println("Solving complexOperations case 4")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	br, bi := ctx.newComplex("b")
	zero := ctx.floatConst(0.0)

	prevCond1 := ar.Eq(zero).And(ai.Eq(zero))
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := br.Eq(zero).And(bi.Eq(zero))
	ctx.Solver.Assert(prevCond2.Not())

	prevCond3 := ar.GT(br)
	ctx.Solver.Assert(prevCond3.Not())

	ctx.Check()
}

func nestedComplexOperations() {
	fmt.Println("Solving nestedComplexOperations")
	nestedComplexOperations1()
	nestedComplexOperations2()
	nestedComplexOperations3()
	nestedComplexOperations4()
}

func nestedComplexOperations1() {
	fmt.Println("Solving nestedComplexOperations case 1")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	zero := ctx.floatConst(0.0)

	cond1 := ar.LT(zero)
	ctx.Solver.Assert(cond1)

	cond2 := ai.LT(zero)
	ctx.Solver.Assert(cond2)

	ctx.Check()
}

func nestedComplexOperations2() {
	fmt.Println("Solving nestedComplexOperations case 2")

	ctx := newContext()
	ar, ai := ctx.newComplex("a")
	zero := ctx.floatConst(0.0)

	cond1 := ar.LT(zero)
	ctx.Solver.Assert(cond1)

	prevCond1 := ai.LT(zero)
	ctx.Solver.Assert(prevCond1.Not())

	ctx.Check()
}

func nestedComplexOperations3() {
	fmt.Println("Solving nestedComplexOperations case 3")

	ctx := newContext()
	ar, _ := ctx.newComplex("a")
	_, bi := ctx.newComplex("b")
	zero := ctx.floatConst(0.0)

	prevCond1 := ar.LT(zero)
	ctx.Solver.Assert(prevCond1.Not())

	cond := bi.LT(zero)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func nestedComplexOperations4() {
	fmt.Println("Solving nestedComplexOperations case 4")

	ctx := newContext()
	ar, _ := ctx.newComplex("a")
	_, bi := ctx.newComplex("b")
	zero := ctx.floatConst(0.0)

	prevCond1 := ar.LT(zero)
	ctx.Solver.Assert(prevCond1.Not())

	prevCond2 := bi.LT(zero)
	ctx.Solver.Assert(prevCond2.Not())

	ctx.Check()
}
