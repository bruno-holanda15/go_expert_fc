package main

func main()  {
	// caso não adicione o buffer, o mesmo terá um deadlock
	ch := make(chan int, 4)

	ch <- 1
	ch <- 2
	ch <- 3

	println(<-ch)
	println(<-ch)
	println(<-ch)
}