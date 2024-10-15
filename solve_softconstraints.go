package main

import "fmt"

func compareAndIncrement() {
	fmt.Println("compareAndIncrement")
	fmt.Println("compareAndIncrement case 1")

	ctx := newContext()
	a := ctx.newInt("a")
	b := ctx.newInt("b")

	cond1 := a.SGT(b)
	ctx.Solver.Assert(cond1)

	c := a.Add(ctx.intConst(1))
	cond2 := c.SGT(b)
	ctx.Solver.Assert(cond2)

	assumptions := []Assumption{
		{a.SLT(ctx.intConst(0)), ctx.Ctx.BoolConst("a < 0")},
		{b.SGT(ctx.intConst(0)), ctx.Ctx.BoolConst("b > 0")},
		{c.SGT(ctx.intConst(0)), ctx.Ctx.BoolConst("c > 0")},
	}

	assumptionsMap := map[string]Assumption{}
	for _, ass := range assumptions {
		assumptionsMap[ass.Name.String()] = ass
	}

	ctx.SolveWithAssumptions(assumptionsMap)
}
