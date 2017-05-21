package main

func main() {
	var readyc chan int

	select {
	case readyc <- 1:
		println("readyc ...")
	default:
		println("default ...")
	}
}
