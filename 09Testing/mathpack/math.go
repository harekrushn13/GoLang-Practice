package mathpack

import "errors"

func Add(a, b int) int {
	return a + b

	//if a > b+10 {
	//	panic("B must be greater than A")
	//}
	//
	//return a + b
}

func FetchUser(id int) (string, error) {
	if id == 1 {
		return "Alice", nil
	}
	return "", errors.New("User not found")
}
