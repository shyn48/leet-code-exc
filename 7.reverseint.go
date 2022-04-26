func reverse(x int) int {
        limit := 2147483648;
	result := 0
	for x != 0 {
		a := x % 10
		result = result * 10 + a
		x = x / 10
	}

    if result > limit || result < -1 * limit {
        return 0
    } else {
        return result
    }
}

