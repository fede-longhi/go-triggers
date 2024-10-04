package utils

func CompareInt(leftValue, rightValue int, operator string) bool {
	switch operator {
	case ">":
		return leftValue > rightValue
	case "<":
		return leftValue < rightValue
	case ">=":
		return leftValue >= rightValue
	case "<=":
		return leftValue <= rightValue
	case "==":
		return leftValue == rightValue
	case "!=":
		return leftValue != rightValue
	default:
		panic("Bad operator type")
	}
}

func CompareFloat(leftValue, rightValue float64, operator string) bool {
	switch operator {
	case ">":
		return leftValue > rightValue
	case "<":
		return leftValue < rightValue
	case ">=":
		return leftValue >= rightValue
	case "<=":
		return leftValue <= rightValue
	case "==":
		return leftValue == rightValue
	case "!=":
		return leftValue != rightValue
	default:
		panic("Bad operator type")
	}
}

func CompareString(leftValue, rightValue string, operator string) bool {
	switch operator {
	case "==":
		return leftValue == rightValue
	case "!=":
		return leftValue != rightValue
	default:
		panic("Bad operator type")
	}
}

func Abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value

}
