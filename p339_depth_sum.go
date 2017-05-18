package goleetcode

func depthSum(num []interface{}) int {
	return depthSumAux(num, 1)
}

func depthSumAux(num []interface{}, w int) int {
	sum := 0
	for _, n := range num {
		switch v := n.(type) {
		case int:
			sum += v * w
		default:
			sum += depthSumAux(v.([]interface{}), w+1)
		}
	}
	return sum
}
