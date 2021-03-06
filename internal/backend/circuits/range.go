package circuits

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gurvy"
)

type rangeCheckConstantCircuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (circuit *rangeCheckConstantCircuit) Define(curveID gurvy.ID, cs *frontend.ConstraintSystem) error {
	c1 := cs.Mul(circuit.X, circuit.Y)
	c2 := cs.Mul(c1, circuit.Y)
	c3 := cs.Add(circuit.X, circuit.Y)
	cs.AssertIsLessOrEqual(c3, 161) // c3 is from a linear expression only
	cs.AssertIsLessOrEqual(c2, 161)
	return nil
}

func rangeCheckConstant() {
	var circuit, good, bad, public rangeCheckConstantCircuit
	r1cs, err := frontend.Compile(gurvy.UNKNOWN, &circuit)
	if err != nil {
		panic(err)
	}

	good.X.Assign(10)
	good.Y.Assign(4)

	bad.X.Assign(11)
	bad.Y.Assign(4)

	public.Y.Assign(4)

	addEntry("range_constant", r1cs, &good, &bad, &public)
}

type rangeCheckCircuit struct {
	X        frontend.Variable
	Y, Bound frontend.Variable `gnark:",public"`
}

func (circuit *rangeCheckCircuit) Define(curveID gurvy.ID, cs *frontend.ConstraintSystem) error {
	c1 := cs.Mul(circuit.X, circuit.Y)
	c2 := cs.Mul(c1, circuit.Y)
	c3 := cs.Add(circuit.X, circuit.Y)
	cs.AssertIsLessOrEqual(c2, circuit.Bound)
	cs.AssertIsLessOrEqual(c3, circuit.Bound) // c3 is from a linear expression only

	return nil
}

func rangeCheck() {

	var circuit, good, bad, public rangeCheckCircuit
	r1cs, err := frontend.Compile(gurvy.UNKNOWN, &circuit)
	if err != nil {
		panic(err)
	}

	good.X.Assign(10)
	good.Y.Assign(4)
	good.Bound.Assign(161)

	bad.X.Assign(11)
	bad.Y.Assign(4)
	bad.Bound.Assign(161)

	public.Y.Assign(4)
	public.Bound.Assign(161)

	addEntry("range", r1cs, &good, &bad, &public)
}

func init() {
	rangeCheckConstant()
	rangeCheck()
}
