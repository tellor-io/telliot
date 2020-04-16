package tracker

type SingleSymbol struct {
	symbol     string
	multiplier float64
	transform  IndexProcessor
}

func (s SingleSymbol)Require() map[string]IndexProcessor {
	r := make(map[string]IndexProcessor)
	r[s.symbol] = s.transform
	return r
}

func (s SingleSymbol)Update(vals map[string]float64) float64 {
	return vals[s.symbol]
}

func (s SingleSymbol)Multiplier() float64 {
	return s.multiplier
}



