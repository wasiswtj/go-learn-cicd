package models

type CalculationResult struct {
	ID         int64
	IntegerOne int64
	IntegerTwo int64
	Result     int64
	Operator   string
}

func (cr *CalculationResult) MapStoreData(p CalcPayload, operator string) {
	cr.IntegerOne = p.IntOne
	cr.IntegerTwo = p.IntTwo
	cr.Result = p.Result
	cr.Operator = operator
}
