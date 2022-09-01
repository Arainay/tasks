package lib

// Дан массив цен за список продуктов, купленных в магазине.
// Товары с одинаковой стоимостью считаются одинаковыми.
// Требуется посчитать, сколько потребуется заплатить суммарно
// за весь поход в магазин при условии, что в магазине
// проводится акция — «купи три одинаковых товара и заплати только за два».
func GetFullCost(costs []int) int {
	var fullCostMap = map[int]int{}
	var fullCost int

	for _, item := range costs {
		fullCostMap[item]++
	}

	for cost, count := range fullCostMap {
		fullCost += cost * (count - count/3)
	}

	return fullCost
}
