package main

import (
	"fmt"
	"main/z3"
)

func arrays() {
	fmt.Println("Solving arrays")
	compareElement()
	compareAge()
}

func compareElement() {
	fmt.Println("Solving compareElement")
	compareElement1()
	compareElement2()
	compareElement3()
	compareElement4()
}

func compareElement1() {
	fmt.Println("Solving compareElement case 1")

	ctx := newContext()
	_, length := ctx.newIntArray("array")
	index := ctx.newInt("index")

	cond := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func compareElement2() {
	fmt.Println("Solving compareElement case 2")

	ctx := newContext()
	array, length := ctx.newIntArray("array")
	index := ctx.newInt("index")
	value := ctx.newInt("value")

	prevCond1 := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(prevCond1.Not())

	element := array.Select(index).(z3.BV)

	cond := element.SGT(value)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func compareElement3() {
	fmt.Println("Solving compareElement case 3")

	ctx := newContext()
	array, length := ctx.newIntArray("array")
	index := ctx.newInt("index")
	value := ctx.newInt("value")

	prevCond1 := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(prevCond1.Not())

	element := array.Select(index).(z3.BV)

	prevCond2 := element.SGT(value)
	ctx.Solver.Assert(prevCond2.Not())

	cond := element.SLT(value)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func compareElement4() {
	fmt.Println("Solving compareElement case 4")

	ctx := newContext()
	array, length := ctx.newIntArray("array")
	index := ctx.newInt("index")
	value := ctx.newInt("value")

	prevCond1 := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(prevCond1.Not())

	element := array.Select(index).(z3.BV)

	prevCond2 := element.SGT(value)
	ctx.Solver.Assert(prevCond2.Not())

	prevCond3 := element.SLT(value)
	ctx.Solver.Assert(prevCond3.Not())

	ctx.Check()
}

func (ctx *Context) personMap() map[string]z3.Sort {
	return map[string]z3.Sort{
		"Name": ctx.Ctx.UninterpretedSort("string"),
		"Age":  ctx.bvSort(),
	}
}
func compareAge() {
	fmt.Println("Solving compareAge")
	compareAge1()
	compareAge2()
	compareAge3()
	compareAge4()
}

func compareAge1() {
	fmt.Println("Solving compareAge case 1")

	ctx := newContext()
	_, length := ctx.newFieldArrays("people", ctx.personMap())
	index := ctx.newInt("index")

	cond := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func compareAge2() {
	fmt.Println("Solving compareAge case 2")

	ctx := newContext()
	people, length := ctx.newFieldArrays("people", ctx.personMap())
	index := ctx.newInt("index")
	value := ctx.newInt("value")

	prevCond1 := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(prevCond1.Not())

	age := people["Age"].Select(index).(z3.BV)

	cond := age.SGT(value)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func compareAge3() {
	fmt.Println("Solving compareAge case 3")

	ctx := newContext()
	people, length := ctx.newFieldArrays("people", ctx.personMap())
	index := ctx.newInt("index")
	value := ctx.newInt("value")

	prevCond1 := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(prevCond1.Not())

	age := people["Age"].Select(index).(z3.BV)

	prevCond2 := age.SGT(value)
	ctx.Solver.Assert(prevCond2.Not())

	cond := age.SLT(value)
	ctx.Solver.Assert(cond)

	ctx.Check()
}

func compareAge4() {
	fmt.Println("Solving compareAge case 4")

	ctx := newContext()
	people, length := ctx.newFieldArrays("people", ctx.personMap())
	index := ctx.newInt("index")
	value := ctx.newInt("value")

	prevCond1 := index.SLT(ctx.intConst(0)).Or(index.SGE(length))
	ctx.Solver.Assert(prevCond1.Not())

	age := people["Age"].Select(index).(z3.BV)

	prevCond2 := age.SGT(value)
	ctx.Solver.Assert(prevCond2.Not())

	cond := age.SLT(value)
	ctx.Solver.Assert(cond)

	ctx.Check()
}
