package utils

import "fmt"

func CatchError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
