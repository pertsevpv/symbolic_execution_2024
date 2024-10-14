package main

func integerOperations(a int, b int) int {
	if a > b {
		return a + b // 1
	} else if a < b {
		return a - b // 2
	} else {
		return a * b // 3
	}
}

func floatOperations(x float64, y float64) float64 {
	if x > y {
		return x / y // 1
	} else if x < y {
		return x * y // 2
	}
	return 0.0 // 3
}

func mixedOperations(a int, b float64) float64 {
	var result float64

	if a%2 == 0 {
		result = float64(a) + b // 1
	} else {
		result = float64(a) - b // 2
	}

	if result < 10 {
		result *= 2 // 1
	} else {
		result /= 2 // 2
	}

	return result
}

func nestedConditions(a int, b float64) float64 {
	if a < 0 {
		if b < 0 {
			return float64(a*-1) + b // 1
		}
		return float64(a*-1) - b // 2
	}
	return float64(a) + b // 3
}

func bitwiseOperations(a int, b int) int {
	if a&1 == 0 && b&1 == 0 {
		return a | b // 1
	} else if a&1 == 1 && b&1 == 1 {
		return a & b // 2
	}
	return a ^ b // 3
}

func advancedBitwise(a int, b int) int {
	if a > b {
		return a << 1 // 1
	} else if a < b {
		return b >> 1 // 2
	}
	return a ^ b // 3
}

func combinedBitwise(a int, b int) int {
	if a&b == 0 {
		return a | b // 1
	} else {
		result := a & b
		if result > 10 {
			return result ^ b // 2
		}
		return result // 3
	}
}

func nestedBitwise(a int, b int) int {
	if a < 0 {
		return -1 // 1
	}

	if b < 0 {
		return a ^ 0 // 2
	}

	if a&b == 0 {
		return a | b // 3
	} else {
		return a & b // 4
	}
}
