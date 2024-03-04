package sprint

func Season(month string) string {
	switch month {
	case "jan", "feb", "dec", "Jan", "Feb", "Dec":
		return "winter"
	case "mar", "apr", "may", "Mar", "Apr", "May":
		return "spring"
	case "jun", "jul", "aug", "Jun", "Jul", "Aug":
		return "summer"
	case "sep", "oct", "nov", "Sep", "Oct", "Nov":
		return "autumn"
	default:
		return "invalid input: " + month
	}
}

