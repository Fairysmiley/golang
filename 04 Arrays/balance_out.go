package sprint

func BalanceOut(arr []bool) []bool {

	var countTrue, countFalse int

	for _, value := range arr {
		if value {
			countTrue++
		} else {
			countFalse++
		}
	}

	for i := 0; countTrue != countFalse; i++ {
		if countTrue < countFalse {
			arr = append(arr, true)
			countTrue++
		} else {
			arr = append(arr, false)
			countFalse++
		}
	}

	return arr
}

