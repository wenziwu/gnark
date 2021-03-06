package circuits

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gurvy"
)

type checkAssertEqualCircuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (circuit *checkAssertEqualCircuit) Define(curveID gurvy.ID, cs *frontend.ConstraintSystem) error {
	cs.AssertIsEqual(circuit.X, circuit.Y)
	c1 := cs.Add(circuit.X, circuit.Y)
	cs.AssertIsEqual(c1, 6)
	return nil
}

func checkAssertEqual() {

	var circuit, good, bad, public checkAssertEqualCircuit
	r1cs, err := frontend.Compile(gurvy.UNKNOWN, &circuit)
	if err != nil {
		panic(err)
	}

	good.X.Assign(3)
	good.Y.Assign(3)

	bad.X.Assign(5)
	bad.Y.Assign(2)

	public.Y.Assign(3)

	addEntry("assert_equal", r1cs, &good, &bad, &public)
}
