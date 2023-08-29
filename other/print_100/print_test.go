package print100

import (
	"fmt"
	"sync"
	"testing"
)

func TestPrint(t *testing.T) {
	printNumber()
}

func print100() {
	a := make(chan int)
	b := make(chan int)
	var wg sync.WaitGroup

	go func(a chan int, b chan int) {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case n := <-a:
				fmt.Printf("a->%d\n", n)
				b <- n + 1
				if n+1 >= 100 {
					return
				}

			}
		}
	}(a, b)
	go func(b chan int, a chan int) {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case n := <-b:

				if n > 100 {
					return
				}
				fmt.Printf("b->%d\n", n)
				a <- n + 1
			}
		}
	}(b, a)
	a <- 1
	wg.Wait()

	close(a)
	close(b)

}

func printOne() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- struct{}{}
			if i%2 == 1 {
				fmt.Println("form-1", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch
			if i%2 == 0 {
				fmt.Println("form-2", i)
			}
		}
	}()
	wg.Wait()
}

// func printTwo() {
// var wg sync.WaitGroup
// var ch1,ch2 =make(chan struct{}),make(chan struct{})
// wg.Add(2)
// go func(){
// 	defer wg.Done()
// 	for i:=1;i<101;i++{
// 		if i%2 ==1{
// 			ftm.
// 		}
// 	}
// }
// }

func abc() {
	a := make(chan bool)
	b := make(chan bool)

	c := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(3)
	f := func(s string, curr <-chan bool, nextch chan bool, wg *sync.WaitGroup) {
		defer wg.Done()
		i := 0
		for i < 10 {
			<-curr
			fmt.Println(s)
			nextch <- true
			i++
		}
	}
	go f("a", a, b, &wg)
	go f("b", b, c, &wg)
	go f("c", c, a, &wg)
	a <- true

	wg.Wait()
}
func TestAbc(t *testing.T) {
	abc()
}
func printNumber() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			ch <- i
			fmt.Println("a->", i)
		}

	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			ch <- i
			fmt.Println("b->", i)
		}

	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	for k := range ch {
		_ = k
	}
}
