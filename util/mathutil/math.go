package mathutil

import (
	"math"

	"advent2024/util/set"
)

func Round(f float64) int {
	return int(math.Round(f))
}

func MinInt(nums ...int) int {
	m := math.MaxInt64
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func MaxInt(nums ...int) int {
	m := math.MinInt64
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func SumOfN(n int) int {
	return n * (n + 1) / 2
}

func SumInts(ints []int) int {
	s := 0
	for _, i := range ints {
		s += i
	}
	return s
}

// angles

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

func RotateDeg(x, y, deg int) (int, int) {
	rad := DegToRad(float64(deg))
	nx, ny := Rotate(float64(x), float64(y), rad)
	return Round(nx), Round(ny)
}

func Rotate(x, y, rad float64) (float64, float64) {
	nx := x*math.Cos(rad) - y*math.Sin(rad)
	ny := x*math.Sin(rad) + y*math.Cos(rad)
	return nx, ny
}

func PrimeFactors(n int) set.Set[int] {
	s := set.NewSet[int]()
	for i := 2; i <= n; i = NextPrime(i) {
		for n%i == 0 {
			s.Add(i)
			n /= i
			if n == 1 {
				break
			}
		}
	}
	return s
}

func IsPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%2 == 0 {
			return false
		}
	}
	return true
}

func NextPrime(n int) int {
	for {
		n++
		if IsPrime(n) {
			return n
		}
	}
}

func IntPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func AbsInt(a int) int {
	return Sign(a) * a
}

func Sign(a int) int {
	if a < 0 {
		return -1
	}
	return 1
}
