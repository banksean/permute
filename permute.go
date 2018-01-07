package permute

// Implements Heap's algorithm for generating permutations.
// Creates the first n permutations of elements in a, and calls f for each.
func permute(n int, a []uint, f func(n []uint)) {
	if n == 1 {
		f(a)
		return
	}

	for i := 0; i < n-1; i++ {
		permute(n-1, a, f)
		if n%2 == 0 {
			a[i], a[n-1] = a[n-1], a[i]
		} else {
			a[0], a[n-1] = a[n-1], a[0]
		}
	}
	permute(n-1, a, f)
}

func permuteCh(n int, a []uint, f chan<- []uint) {
	if n == 1 {
		f <- a
		return
	}

	for i := 0; i < n-1; i++ {
		permuteCh(n-1, a, f)
		if n%2 == 0 {
			a[i], a[n-1] = a[n-1], a[i]
		} else {
			a[0], a[n-1] = a[n-1], a[0]
		}
	}
	permuteCh(n-1, a, f)
}

// Call f with a permutation mapping for each permutation of n distinct elements.
func permuteN(n int, f func(n []uint)) {
	a := make([]uint, n)
	for i := 0; i < n; i++ {
		a[i] = uint(i)
	}
	permute(n, a, f)
}

func permuteNCh(n int) chan []uint {
	f := make(chan []uint)
	go func() {
		a := make([]uint, n)
		for i := 0; i < n; i++ {
			a[i] = uint(i)
		}
		permuteCh(n, a, f)
		close(f)
	}()
	return f

