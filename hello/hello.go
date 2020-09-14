package hello

// import (
// "fmt"
// )

func Reserve(str string) string {
	// fmt.Println(global)
	reserveStr := []rune(str)

	for i, j := 0, len(str)-1; i < len(str)/2; j, i = j-1, i+1 {
		reserveStr[i], reserveStr[j] = reserveStr[j], reserveStr[i]
	}

	return string(reserveStr)
}
