package factorial

//Recur computes the factorial of n with recursion calls
func Recur(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Recur(n-1)
}

//Iter computes the factorial of n with basic for loop
func Iter(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
