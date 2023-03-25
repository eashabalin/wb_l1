package task26

import "fmt"

func uniqueSymbols(s string) bool {
	m := make(map[int32]int)

	for _, symbol := range s {
		m[symbol] += 1
		if v, ok := m[symbol]; ok && v > 1 {
			return false
		}
	}
	return true
}

func Run() {
	fmt.Println("--- Task 26 ---")

	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(s1, uniqueSymbols(s1))
	fmt.Println(s2, uniqueSymbols(s2))
	fmt.Println(s3, uniqueSymbols(s3))

}
