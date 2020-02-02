package arrayext

//求并集
func union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//求最多值 [1 7 10 4 8 6 9 10 2 43]  //9 43
func GetArrayMaxValue(intSlice []int) (int, int) {
	maxVal := intSlice[0]
	maxValIndex := 0
	for i := 0; i < len(intSlice); i++ {
		//从第二个元素开始循环比较，如果发现有更大的数，则交换
		if maxVal < intSlice[i] {
			maxVal = intSlice[i]
			maxValIndex = i
		}
	}
	return maxValIndex, maxVal
}
