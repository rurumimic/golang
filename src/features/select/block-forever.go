package main

func main() {
	select {}
}

/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select (no cases)]:
*/
