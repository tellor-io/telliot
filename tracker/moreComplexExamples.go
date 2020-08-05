package tracker

// ChainedPrice computes a pair that doesn't exist by using multiple symbols
type ChainedPrice struct {
	chain       []string
	transform   IndexProcessor
	granularity float64
}

//func (c ChainedPrice) Require(at time.Time) map[string]IndexProcessor {
//	r := make(map[string]IndexProcessor)
//	for _, symbol := range c.chain {
//		r[symbol] = c.transform
//	}
//	return r
//}

//func (c ChainedPrice) ValueAt(vals map[string]float64, at time.Time) float64 {
//	val := 1.0
//	for _, pair := range c.chain {
//		val *= vals[pair]
//	}
//	fmt.Println("Ample Chained Price ", val*c.granularity)
//	return val * c.granularity
//}
