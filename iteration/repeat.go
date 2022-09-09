package iteration

func Repeat(repeatChar string, times int) (text string) {
	for count := 0; count < times; count++ {
		text += repeatChar
	}
	return
}
