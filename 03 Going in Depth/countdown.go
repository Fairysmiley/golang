package sprint

func Countdown(n int) string {

	if n < 0 {
		return "Invalid input"
	}
	
	var result string
	for i := n; i > 0; i -= 2 {
		result += (string('0' + i) + ", ")
	}
	return result + "0!"
}

