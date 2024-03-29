package main

import "fmt"

func main() {

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for { // keep repeating
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (_, _ <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			for val := range orDone(done, in) { // 1 2 1 2
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ { // 0 1 <- to ensure both are written
					select {
					case <-done:
					case out1 <- val:
						out1 = nil // block writing
					case out2 <- val:
						out2 = nil // block writing
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan interface{})
	defer close(done)

	out1, out2 := tee(done, take(done, repeat(done, 1, 2), 4)) // [1 2 1 2] 1 2 1 2 1 2...

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}

}

/*
out1: 1, out2: 1
out1: 2, out2: 2
out1: 1, out2: 1
out1: 2, out2: 2
*/
