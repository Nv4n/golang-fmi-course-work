package sqrt_utils

import "math"

func Newton(num float64) (z float64, iter int) {
	z, iter, epsilon := num, 0, 1e-10
	for {
		iter++
		root := 0.5 * (z + (num / z))
		if math.Abs(root-z) < epsilon {
			return
		}
		z = root
	}
}
