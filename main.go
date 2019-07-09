package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	. "github.com/spotmaxtech/spot-consul/spotconsul"
	"os"
	"os/signal"
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

	// run the logic
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		logic.RunningLoop(ctx)
		wg.Done()
	}(ctx)

	// cancel context controlling
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		log.Warnf("received an interrupt, stopping service (%d sec) ...", config.Global.LingerTimeS)
		cancel()
		time.Sleep(time.Second * time.Duration(config.Global.LingerTimeS))
		wg.Done()
	}()
	wg.Wait()
	log.Infof("service stopped")
}
