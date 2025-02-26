package main

func main() {

}

//func fanIn(inputs ...<-chan int) <-chan int {
//	var wg sync.WaitGroup
//	out := make(chan int)
//
//}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, value := range work {
			ch <- value
		}
	}()

	return ch
}
