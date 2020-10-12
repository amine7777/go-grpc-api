package functions

func AddFunc(a int64, b int64) int64 {
	return a + b
}

func MultFunc(a int64, b int64) int64 {
	return a * b
}

func PrimeFunc(number int64) (r bool) {
	if number > 1 {
		var i int64
		for i = 2; i <= number; i++ {
			if number%2 == 0 {
				r = true
				break

			} else {
				r = false
				break
			}
		}

	} else {
		r = false
	}
	return r
}
