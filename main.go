package main

func main() {

}

func incomeOfWeek(incomes [] int) (sum int) {
	for _, income := range incomes {
		sum += income
	}
	return sum
}

func compareOfWeekIncome(sum1, sum2, sum3 int) (res string) {
	max := 0
	if max <= sum1{
		max = sum1
		if max <= sum2{
			max = sum2
			if max <= sum3{
				max = sum3
				return "3"
			}
			return "2"
		}
		return "1"
	}
	return "0"
}