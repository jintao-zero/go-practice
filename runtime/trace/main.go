package main

import (
	"os"
	"context"
	"log"
	"runtime/trace"
	"time"
)

func main()  {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()
 	makeCoffee()
	makeCoffee()

	time.Sleep(3 * time.Second)
}

func makeCoffee()  {
	ctx, task := trace.NewTask(context.Background(), "makeCappuccino")
	trace.Log(ctx, "orderID", "11111")

	milk := make(chan bool)
	espresso := make(chan bool)

	go func() {
		trace.WithRegion(ctx, "steamMilk", func() {
			time.Sleep(time.Second)
		})
		milk <- true
	}()
	go func() {
		trace.WithRegion(ctx, "extractCoffee", func() {
			time.Sleep(time.Second)
		})
		espresso <- true
	}()
	go func() {
		defer task.End() // When assemble is done, the order is complete.
		<-espresso
		<-milk
		trace.WithRegion(ctx, "mixMilkCoffee", func() {
			time.Sleep(time.Second)
		})
	}()
	time.Sleep(time.Second)
}