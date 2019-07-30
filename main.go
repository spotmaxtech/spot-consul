package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"github.com/spotmaxtech/gokit"
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
	// log init with hook
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetURL("http://es.spotmaxtech.com"),
		elastic.SetBasicAuth("", ""),
		elastic.SetSniff(false))
	hook, err := gokit.NewElasticHook(client, host, log.DebugLevel, "spot-consul", gokit.RotateMonth)
	if err != nil {
		panic(err)
	}
	log.AddHook(hook)
	log.SetLevel(log.DebugLevel)
	config := NewConfig("./configs/regions.json")

	fmt.Println(Prettify(config))

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for _, logicCfg := range config.Logic {
		logic := NewLearningLogic(logicCfg, config.Global)
		log.Infof("ready to run service %#v", logicCfg)
		wg.Add(1)
		go func(ctx context.Context) {
			logic.RunningLoop(ctx) // here logic will use ctx too
			wg.Done()
		}(ctx)
	}

	// cancel context controlling
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		log.Warnf("received an interrupt, stopping service (%d sec) ...", config.Global.LingerTimeS)
		cancel()
		time.Sleep(time.Second * time.Duration(config.Global.LingerTimeS))
	}()
	wg.Wait()
	log.Infof("service stopped")
}
