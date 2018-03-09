package main

type Operator interface {
	Apply(map[string]map[string]string)
}

type StrategyOperation struct {
	Operator Operator
}

func (o *StrategyOperation) Operate(container map[string]map[string]string) {
	o.Operator.Apply(container)
}
