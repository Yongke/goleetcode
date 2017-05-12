package goleetcode

func distributeCandies(candies []int) int {
	l := len(candies)
	m := map[int]int{}
	for i:=0;i<l;i++{
		m[candies[i]] = 0
	}
	l1 := len(m)
	if l1 >= l/2 {
		return l/2
	}
	return l1
}
