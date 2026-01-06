package buildstronghold

import "fmt"

const MOD = 1000000007

func PowExpDemo() {
	res := PowIter(2, 60)

	fmt.Println("res is : ", res)
}

// implement int first, x = base, n = exponent
// recursive impl.
func PowRec(x, n int) int {
	if n == 0 {
		return 1
	}

	// for n != 0 {
	var result int
	// IMP : bit opertion seems to be more efficient.
	if n&1 == 1 {
		// odd case
		// update exp and result
		// n = n - 1
		// result = (result * x) % MOD
		result = PowRec(x, n-1) * x
		// fmt.Println("result - o", result)
	} else {
		// even case
		k := n / 2
		y := PowRec(x, k)
		// result = ((y % MOD) * (y % MOD)) % MOD
		result = y * y
	}
	// }

	return result
}

//TODO - MV : write modular exponentiation

// iterative impl.
func PowIter(x, n int) int {
	if n == 0 {
		return 1
	}

	result := 1
	for n != 0 {
		if n&1 == 1 {
			// odd case
			// update exp and result
			n = n - 1
			result = (result * x)
			// fmt.Println("result - o", result)
		} else {
			// even case

			// with below we are able to achieve halfing the exponent and
			// and kind of sharing the workload by squaring the base itself.
			n = n / 2
			// update the base
			x = x * x
		}
	}

	return result
}
