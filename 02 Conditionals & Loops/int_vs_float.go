package sprint

func IntVsFloat(i int, f float32) string {
	if float64(f) < float64(i) {
		return "Integer"
	} else if float64(f) > float64(i) {
		return "Float"
	} else {
		return "Same"
	}
}
