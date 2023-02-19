package utils

func CatchError(err error) {
	if err != nil {
		panic(err)
	}
}
