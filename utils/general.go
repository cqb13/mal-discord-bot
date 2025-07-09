package utils

func Ternary(cond bool, a string, b string) string {
	if cond {
		return a
	}
	return b
}
