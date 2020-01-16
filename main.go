package main

func main() {

}

func incomeOfWeek(incomes [] int) (sum int) {
	for _, income := range incomes {
		sum += income
	}
	return sum
}

func compareOfWeekIncome(sum1, sum2, sum3 int) (result string) {
	max := 0
	if max <= sum1 {
		max = sum1
		if max <= sum2 {
			max = sum2
			if max <= sum3 {
				max = sum3
				return "3"
			}
			return "2"
		}
		return "1"
	}
	return "0"
}

func compareTheBestDailyMiddleIncome(sum1, sum2, sum3 int) (result string) {
	sum1 /= 7
	sum2 /= 7
	sum3 /= 7
	maxIncome := 0
	if maxIncome <= sum1 {
		maxIncome = sum1
		if maxIncome <= sum2 {
			maxIncome = sum2
			if maxIncome <= sum3 {
				maxIncome = sum3
				return "3"
			}
			return "2"
		}
		return "1"
	}
	return "0"
}
