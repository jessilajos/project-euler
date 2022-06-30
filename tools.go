package problem

func arithmeticMedian(n, k int) int {
	return int((n + k) / 2.0)
}

func arithmeticSum(n, k int) int {
	remainder := 0
	for n%k != 0 {
		n++
		remainder++
	}
	if remainder > 0 {
		return (n*(n+k)/k)/2.0 - n
	}
	return (n * (n + k) / k) / 2.0
}
