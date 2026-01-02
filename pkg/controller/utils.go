package controller

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
