package fib

var (
	cache map[int]int
)

func init() {
	cache = make(map[int]int)
}

// Fib returns a fibonacci number converted from an int argument
func Fib(num int) int {
	if num <= 1 {
		return num
	}

	if result, ok := cache[num]; ok {
		return result
	}

	sum := Fib(num-1) + Fib(num-2)
	cache[num] = sum

	return sum
}
