package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	. "github.com/spotmaxtech/spot-consul/spotconsul"
	"sync"
	"time"
)

func Process(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			default:
				fmt.Println("sleep")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	wg.Wait()
}

func main() {
	config := NewConfig("./configs/regions.json")
	log.SetLevel(log.DebugLevel)
	fmt.Println(Prettify(config))

	logic := NewLearningLogic(config.Logic[0], config.Global)

	ctx, _ := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		logic.RunningLoop(ctx)
		wg.Done()
	}(ctx)
	wg.Wait()
}
