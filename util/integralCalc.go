package util

import "math"

var nums [10]rune = [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

type term struct {
	Degree int64
	Number float64
	PreOp  rune
}

func evaluate(f []term, x float64) float64 {
	value := 0.
	for _, term := range f {
		subval := math.Pow(term.Number*x, float64(term.Degree))
		if term.PreOp == '+' {
			value += subval
		} else if term.PreOp == '-' {
			value -= subval
		} else if term.PreOp == '/' {
			value /= subval
		} else if term.PreOp == '*' {
			value *= subval
		}
	}

	return value
}

func Integral(lowerBound uint, upperBound uint, f string, diffVar rune, delta float64) float64 {
	var eval []term
	eval = make([]term, 0)
	coeff := 0.
	digitCount := 0
	state := 1
	inTerm := true
	preop := '+'
	finalValue := 0.
	for _, char := range f {
		value, valid := getNum(char)

		if valid && inTerm {
			coeff += float64(value) * math.Pow(10, float64(digitCount))
			digitCount += 1 * state
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			preop = char
		} else if char == '.' {
			digitCount = 1
			state = -1
		} else if diffVar == char {
			inTerm = false
		} else if valid {
			eval = append(eval, term{PreOp: preop, Degree: int64(value), Number: float64(coeff)})
			coeff = 0.
			digitCount = 0
			state = 1
			inTerm = true
		}
	}

	for i := float64(lowerBound); i <= float64(upperBound); i += delta {
		finalValue += delta * evaluate(eval, i)
	}

	return finalValue
}

func getNum(tok rune) (int, bool) {
	for idx, num := range nums {
		if tok == num {
			return idx, true
		}
	}

	return -1, false
}
