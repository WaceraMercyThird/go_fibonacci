package app

// Iterative Algorithm
func FibonacciIteration(n int) int {

	if n <= 1 {

		return n
	}

	var n2, n1 = 0, 1

	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	
	return n1
}

// Recursive Algorithm
func FibonacciRecursive(n int) int {

	if n <= 1 {

		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// Dynamic Programming Algorithm
func FibonacciDP(n int) int {

	if n <= 0 {

		return n
	}

	f := make([]int, n+2)

	f[0], f[1] = 0, 1

	for i := 2; i <= n; i++ {

		f[i] = f[i-1] + f[i-2]

	}

	return f[n]

}
