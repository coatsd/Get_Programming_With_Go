package main

import (
	"fmt"
	"strings"
	"sync"
)

type Info struct {
	routine int
	node string
}

func main() {
	data := make(chan string)
	go func() {
		defer close(data)
		for _, s := range fetch_data() {
			data <- s
		}
	}()

	infos := make(chan *Info)
	go func() {
		defer close(infos)
		var wg sync.WaitGroup
		add_data_jobs := 2
		wg.Add(add_data_jobs)

		for iter := 0; iter < add_data_jobs; iter++ {
			go func(id int) {
				defer wg.Done()
				for d := range data {
					i := &Info {
						routine: id,
						node: strings.ToUpper(d),
					}
					infos <- i
				}
			}(iter)
		}
		fmt.Println("waiting")
		wg.Wait()
		fmt.Println("done")
	}()

	for i := range infos {
		fmt.Println(i)
	}
}

func fetch_data() []string {
	return []string {
		"Thing 1",
		"Thing 2",
		"thing 3",
		"Thing 4",
		"Thing 5",
		"Thing 6",
		"Thing 7",
		"Thing 8",
	}
}