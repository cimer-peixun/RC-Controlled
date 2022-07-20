package shared

func Panic(err error) {
	if BUILDVARIABLE != BUILDRELEASE {
		panic(err)
	}
}
