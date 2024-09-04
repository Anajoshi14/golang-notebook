package phone

import (
	C "github.com/Anajoshi14/golang_repo/calculator"
)

type Phone struct {
	Calci C.Calculator
}

func CreatePhone(cal C.Calculator) Phone {
	return Phone{
		Calci: cal,
	}
}
