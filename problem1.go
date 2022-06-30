package projecteuler

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.

func jessisolutionP1(input interface{}) (string, interface{}) {
	n := input.(int)
	sum3 := 0
	sum5 := 0
	sum15 := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 {
			sum3 += i
		}
		if i%5 == 0 {
			sum5 += i
		}
		if i%15 == 0 {
			sum15 += i
		}
	}
	return "jessi", sum3 + sum5 - sum15
}

func prestonsolutionP1(input interface{}) (string, interface{}) {
	n := input.(int)
	sum := 0
	for i := 3; i < n; i += 3 {
		sum += i
	}
	for i := 5; i < n; i += 5 {
		if i%15 != 0 {
			sum += i
		}
	}
	return "preston", sum
}

func gracefulArithmeticSumP1(n, k int) int {
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

func jointsolutionP1(input interface{}) (string, interface{}) {
	n := input.(int) - 1
	sum3 := gracefulArithmeticSumP1(n, 3)
	sum5 := gracefulArithmeticSumP1(n, 5)
	sum15 := gracefulArithmeticSumP1(n, 15)
	return "joint", sum3 + sum5 - sum15
}
