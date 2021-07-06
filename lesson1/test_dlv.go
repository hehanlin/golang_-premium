package main

func main() {
	var ch chan int
	send(ch)
}

func send(ch chan int) {
	ch <- 1
}
