package channel

func Process1(c chan string, data string) {
	c <- data
}