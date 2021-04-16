package main

import (
	"context"
	"fmt"
	"time"
)

type WithContext func(context.Context, string) (string, error)

type SlowFunction func(string) (string, error)

func Timeout(f SlowFunction) WithContext {
	return func(ctx context.Context, arg string) (string, error) {
		chres := make(chan string)
		cherr := make(chan error)

		go func() {
			res, err := f(arg)
			chres <- res
			cherr <- err
		}()

		select {
		case res := <-chres:
			return res, <-cherr
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func Slow(x string) (string, error) {
	// Simulating slowness/delay...
	time.Sleep(2 * time.Second)
	return "OK...", nil
}

func main() {
	ctx := context.Background()
	ctxt, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	timeout := Timeout(Slow)
	res, err := timeout(ctxt, "some input")
	if err != nil {
		fmt.Println("There was an error ... :(")
		fmt.Println(err)
	} else {
		fmt.Printf("Everything OK -> [%s]\n", res)
	}
}
