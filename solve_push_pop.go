package main

import "fmt"

func pushPopIncrementality() {
	fmt.Println("Solving pushPopIncrementality")

	ctx := newContext()
	j := ctx.newInt("j")
	result := j
	for i := int64(1); i <= 10; i++ {
		result = result.Add(ctx.intConst(i))
	}

	cond := result.SMod(ctx.intConst(2)).Eq(ctx.intConst(0))

	ctx.Solver.Push()
	ctx.Solver.Assert(cond)
	ctx.Check()
	ctx.Solver.Pop()

	ctx.Solver.Assert(cond.Not())
	ctx.Check()
}
