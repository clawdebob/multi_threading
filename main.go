package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func thread(n int, i *int, typ int, can chan int) {
	dur := time.Duration(rand.Intn(5000))
	time.Sleep(time.Millisecond * dur)
	fmt.Println("Поток:", n, "проспал:", int(dur), "миллисекунд")
	mutex.Lock()
	switch typ {
	case 0:
		for c := 0; c < 10; c++ {
			*i++
			time.Sleep(time.Millisecond * 150)
		}
	case 1:
		for c := 0; c < 20; c++ {
			*i += 5
			time.Sleep(time.Millisecond * 200)
		}
	case 2:
		*i++
		*i = -*i
	case 3:
		*i--
		*i *= 5
	case 4:
		for c := 0; c < 50; c++ {
			*i -= 4
			time.Sleep(time.Millisecond * 100)
		}
	case 5:
		*i *= 10
		*i++
	}
	fmt.Println("Поток:", n, "выполнял действие номер:", typ, "результат:", *i)
	mutex.Unlock()
	can <- n
}

func main() {
	var (
		change int = 0
		number int
	)
	c := make(chan int)
	fmt.Printf("Введите количество потоков: ")
	fmt.Scanf("%d", &number)
	ticker := time.Now()
	for i := 1; i <= number; i++ {
		go thread(i, &change, rand.Intn(6), c)
	}
	for i := 0; i < number; i++ {
		<-c
	}
	fmt.Println(time.Since(ticker))

}
