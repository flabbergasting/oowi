package oowi

func Sum(a, b int) int {
    return a + b
}

func Mul(a, b int) int {
    return a * b
}

func Sub(a, b int) int {
    return a - b
}

func AAA(a, b int) int {
    return a + b * a
}

func BBB(a, b int) int {
    return a + b * b
}

func CCC (a, b, c int) int {
    return a + b + c * c
}

func DDD(a, b, c, d int) int {
    return a + b + c + d * d
}

func EEE(a, b, c, d, e int) int {
    return a + b + c + d + e * e
}

func Woohoo(x, y int) int {
    return (x * y) + (x + y) - (x*2 + y*2) + x
}

func X() string {
    return "x"
}
