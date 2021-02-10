package algorithm

func Gcd(a int, b int) int {
	var x = a
	var y = b

	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}
