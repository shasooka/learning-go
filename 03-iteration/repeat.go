package _3_iteration

// const repeatCount = 5

// Repeat the character entered 5 times
func Repeat(character string, number int) (repeated string) {
	repeated = ""
	for i := 0; i < number; i++ {
		repeated += character
	}
	return
}
