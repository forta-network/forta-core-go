package inspect

// Results
const (
	ResultSuccess         float64 = 1
	ResultFailure         float64 = -1
	ResultInternalProblem float64 = -2
	ResultUnknown         float64 = -3
)

// defaultIndicators returns a map with all indicators set to ResultUnknown
func defaultIndicators(indicatorKeys []string) (indicators map[string]float64) {
	indicators = make(map[string]float64)

	for _, indicator := range indicatorKeys {
		indicators[indicator] = ResultUnknown
	}

	return
}
