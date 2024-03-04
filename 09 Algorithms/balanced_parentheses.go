package sprint

func BalancedParentheses(input string) bool {

	squared := 0
	rounded := 0
	curly := 0

	for _, char := range input {
		switch char {
		case '[':
			squared++
		case ']':
			squared--
		case '(':
			rounded++
		case ')':
			rounded--
		case '{':
			curly++
		case '}':
			curly--
		}
	}
	if rounded == 0 && squared == 0 && curly == 0 {
		return true
	}
	return false
}
