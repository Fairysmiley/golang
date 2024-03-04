package sprint

func LongestCommonSubstr(str1, str2 string) string {

	result := ""

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				k := 0
				for k+i < len(str1) && k+j < len(str2) && str1[i+k] == str2[j+k] {
					k++
				}
				if k > len(result) {
					result = str1[i : i+k]
				}
			}
		}
	}
	return result
}
