package helper

//Month helper for get number of month
func Month(m string) int {
	var monthNumber int
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	for i := 0; i < len(months); i++ {
		if months[i] == m {
			monthNumber = i + 1
		}
	}

	if monthNumber > 0 {
		return monthNumber
	}

	return 0
}
