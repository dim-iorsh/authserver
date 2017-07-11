package main

import (
	//"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Mechanical domain.
	errc := make(chan error)
	//ctx := context.Background()

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

}
