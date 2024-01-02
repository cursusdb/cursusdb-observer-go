# CursusDB Observer for GO

## About Node Observers
https://cursusdb.com/documentation#realtime-observers

### Example
``` 
package main

import (
	"fmt"
	observer "github.com/cursusdb/cursusdb-observer-go"
)

func main() {
	wg := &sync.WaitGroup{}
	var sigCh chan os.Signal
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	var ob *observer.Observer

	ob = &observer.Observer{
		Host:      "0.0.0.0",
		Port:      7680,
		SharedKey: "yoursharedkey",
		Channel:   make(chan string),
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := ob.Listen()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case sig := <-sigCh:
				log.Println("received signal", sig)
				return
			case ev := <-ob.Channel:
				log.Println("New node event: ", ev)
			default:
				time.Sleep(time.Nanosecond * 1000000)
				continue
			}
		}
	}()
	
	wg.Wait()

	defer ob.Close()
}
```