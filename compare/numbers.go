package compare

//minInt of two integers
func minInt(a int, b int) (res int) {
	if a < b {
		return a
	} else {
		return b
	}
}

//max of two integers
func maxInt(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

//max of two float64s
func max(a float64, b float64) float64 {
	if a < b {
		return b
	} else {
		return a
	}
}

//min of two float64s
func min(a float64, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}
