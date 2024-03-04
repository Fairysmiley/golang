package sprint

func Payout(amount int, denominations []int) (payout []int) {

	nominal := len(denominations) - 1

	denominations = Sort_Arr(denominations)

	for amount != 0 {
		if nominal == 0 && amount-denominations[nominal] < 0 {
			return []int{}
		}
		if amount >= denominations[nominal] {
			payout = append(payout, denominations[nominal])
			amount -= denominations[nominal]
		} else {
			nominal--
		}

	}
	return payout

}

func Sort_Arr(a []int) []int {

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a

}
