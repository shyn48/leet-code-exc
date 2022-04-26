func reverse(x int) int {
    limit := 2147483648;

	result := 0
	power := 0
	temp := x
	for temp != 0 {
		println(temp)
		temp = temp / 10
		power++
	}

	power -= 1

	for x != 0 {
		a := x % 10
		fmt.Println(a, power, result)
		result = result + int((float64(a) * math.Pow(10, float64(power))))
		power -= 1
		x = x / 10
	}

    if result > limit || result < -1 * limit {
        return 0
    } else {
        return result
    }
}

