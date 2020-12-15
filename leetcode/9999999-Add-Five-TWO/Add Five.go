package leetcode

func addFive(n string) string {
	index := 0
	if n[0] >= '0' && n[0] <= '9' {
		for index = 0; index < len(n); index++ {
			if int(n[index]-'0') < 5 {
				break
			}
		}
		return n[:index] + "5" + n[index:]
	}
	for index = 1; index < len(n); index++ {
		if int(n[index]-'0') > 5 {
			break
		}
	}
	return n[:index] + "5" + n[index:]
}
