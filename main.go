package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	creational "github.com/ovijeet26/golang-concepts/golang-concepts/designPatterns/creational"

	contextLogic "github.com/ovijeet26/golang-concepts/golang-concepts/context"
)

func main() {

	creational.FactoryClientCode()
	creational.BuilderClientCaller()
	designPatternCaller()
	contextCaller()
}

func contextCaller() {
	start := time.Now()

	ctx := context.Background()

	fmt.Print(contextLogic.MockCaller(ctx))

	fmt.Print("Total time taken -> ", time.Since(start))
}

func designPatternCaller() {

	singleonCaller()
}

func singleonCaller() {

	wg := sync.WaitGroup{}
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go threadCaller(&wg)
	}
	fmt.Println("Singleton call ends")
}

func threadCaller(wg *sync.WaitGroup) {

	wg.Done()
	creational.GetInstance()
}
