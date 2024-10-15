package main

import (
	"fmt"
	"main/z3"
)
import "math"

var bits = 32

var maxInt = int64(math.MaxInt32)
var minInt = int64(math.MinInt32)

var maxFloat = math.MaxFloat64
var minFloat = -maxFloat

type Context struct {
	Solver *z3.Solver
	Ctx    *z3.Context
}

func newContext() Context {
	ctx := z3.NewContext(&z3.Config{})
	solver := z3.NewSolver(ctx)
	ctx.Config().SetBool("unsat_core", true)
	return Context{
		Solver: solver,
		Ctx:    ctx,
	}
}

func (ctx *Context) Check() {
	fmt.Println(ctx.Solver.String())

	sat, err := ctx.Solver.Check()
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return
	}

	fmt.Println("Sat: ", sat)
	if !sat {
		unsatCore := ctx.Solver.GetUnsatCore()
		for i := range unsatCore {
			fmt.Println(unsatCore[i])
		}
	} else {
		fmt.Println(ctx.Solver.Model().String())
	}
	fmt.Println("===============")
}

func (ctx *Context) bvSort() z3.Sort {
	return ctx.Ctx.BVSort(32)
}

func (ctx *Context) intConst(value int64) z3.BV {
	return ctx.Ctx.FromInt(value, ctx.bvSort()).(z3.BV)
}

//func (ctx *Context) minInt() z3.Int {
//	return ctx.intConst(minInt)
//}
//
//func (ctx *Context) maxInt() z3.Int {
//	return ctx.intConst(maxInt)
//}

func (ctx *Context) newInt(name string) z3.BV {
	newInt := ctx.Ctx.BVConst(name, bits)
	//ctx.Solver.Assert(newInt.GE(ctx.minInt()).And(newInt.LE(ctx.maxInt())))
	return newInt
}

func (ctx *Context) float64Sort() z3.Sort {
	return ctx.Ctx.FloatSort(11, 53)
}

func (ctx *Context) floatConst(value float64) z3.Float {
	return ctx.Ctx.FromFloat64(value, ctx.float64Sort())
}

func (ctx *Context) minFloat() z3.Float {
	return ctx.floatConst(minFloat)
}

func (ctx *Context) maxFloat() z3.Float {
	return ctx.floatConst(maxFloat)
}

func (ctx *Context) float(value z3.BV) z3.Float {
	return value.SToFloat(ctx.float64Sort())
}

func (ctx *Context) int(value z3.Float) z3.Int {
	return value.ToReal().ToInt()
}

func (ctx *Context) newFloat(name string) z3.Float {
	newFloat := ctx.Ctx.Const(name, ctx.float64Sort()).(z3.Float)
	ctx.Solver.Assert(newFloat.GE(ctx.minFloat()).And(newFloat.LE(ctx.maxFloat())))
	return newFloat
}

// return real, imag
func (ctx *Context) newComplex(name string) (z3.Float, z3.Float) {
	return ctx.newFloat(name + ".real"), ctx.newFloat(name + ".imag")
}

func (ctx *Context) intArraySort() z3.Sort {
	return ctx.Ctx.ArraySort(ctx.Ctx.BVSort(bits), ctx.Ctx.BVSort(bits))
}

// return array, array.length
func (ctx *Context) newIntArray(name string) (z3.Array, z3.BV) {
	length := ctx.newInt(name + ".length")
	ctx.Solver.Assert(length.SGE(ctx.intConst(0)))

	array := ctx.Ctx.Const(name+".array", ctx.intArraySort()).(z3.Array)
	return array, length
}

// return array, array.length
func (ctx *Context) newFieldArrays(name string, sorts map[string]z3.Sort) (map[string]z3.Array, z3.BV) {
	length := ctx.newInt(name + ".length")
	ctx.Solver.Assert(length.SGE(ctx.intConst(0)))

	fieldArrays := map[string]z3.Array{}

	for field, sort := range sorts {
		fieldArraySort := ctx.Ctx.ArraySort(ctx.bvSort(), sort)
		fieldArrays[field] = ctx.Ctx.Const(name+"."+field+".array", fieldArraySort).(z3.Array)
	}

	return fieldArrays, length
}
