package task10

import "fmt"

func min(arr []float64) (val float64, index int) {
	mIndex := 0
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < m {
			m = arr[i]
			mIndex = i
		}
	}
	return m, mIndex
}

func Run() {
	fmt.Println("--- Task 10 ---")

	values := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, v := range values {
		group := int(v/10) * 10
		groups[group] = append(groups[group], v)
	}

	fmt.Println(groups)

}
