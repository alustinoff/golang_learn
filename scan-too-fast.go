package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p:=range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main(){
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i:=0; i<cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i:=1; i<=1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

/*func main(){
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}*/