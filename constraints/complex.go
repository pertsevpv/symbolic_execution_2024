package main

func basicComplexOperations(a complex128, b complex128) complex128 {
	if real(a) > real(b) {
		return a + b // 1
	} else if imag(a) > imag(b) {
		return a - b // 2
	}
	return a * b // 3
}

func complexMagnitude(a complex128) float64 {
	magnitude := real(a)*real(a) + imag(a)*imag(a)
	return magnitude
}

func complexComparison(a complex128, b complex128) string {
	magA := complexMagnitude(a)
	magB := complexMagnitude(b)

	if magA > magB {
		return "Magnitude of a is greater than b" // 1
	} else if magA < magB {
		return "Magnitude of b is greater than a" // 2
	}
	return "Magnitudes are equal" // 3
}

func complexOperations(a complex128, b complex128) complex128 {
	if real(a) == 0 && imag(a) == 0 {
		return b // 1
	} else if real(b) == 0 && imag(b) == 0 {
		return a // 2
	} else if real(a) > real(b) {
		return a / b // 3
	}
	return a + b // 4
}

func nestedComplexOperations(a complex128, b complex128) complex128 {
	if real(a) < 0 {
		if imag(a) < 0 {
			return a * b // 1
		}
		return a + b // 2
	}

	if imag(b) < 0 {
		return a - b // 3
	}
	return a + b // 4
}
