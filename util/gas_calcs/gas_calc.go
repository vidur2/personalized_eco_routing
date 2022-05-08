package gascalc

func CalculateMpg(defaultMpg float32, speed float32) float32 {
	if speed > 50 {
		return defaultMpg - ((speed - 50) / 5)
	} else {
		return speed
	}
}
