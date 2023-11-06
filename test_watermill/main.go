package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var ErrRouterIsNotRunning error = errors.New("Router is not running")

type Queue struct {
	mapRouter map[string]*Router
	locker    sync.Mutex
}

type Router struct {
	router *message.Router
	pub    message.Publisher
	sub    message.Subscriber
}

func (q Queue) AddRouter(name string) error {
	q.locker.Lock()
	defer q.locker.Unlock()
	if q.mapRouter[name] != nil {
		return nil
	} else {
		router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
		if err != nil {
			return err
		}
		pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
		q.mapRouter[name] = &Router{
			router: router,
			pub:    pubSub,
			sub:    pubSub,
		}
	}
	router := q.mapRouter[name]
	router.router.AddNoPublisherHandler(name, name, router.sub, q.HandlerRequest)
	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	ctx := context.Background()
	go router.router.Run(ctx)
	<-router.router.Running()
	return nil
}

func (q Queue) SendMessage(name string, msg *message.Message) error {
	if q.mapRouter[name] == nil {
		err := q.AddRouter(name)
		if err != nil {
			return err
		}
	}
	if q.mapRouter[name].router.IsRunning() {
		err := q.mapRouter[name].pub.Publish(name, msg)
		if err != nil {
			return err
		}
		return nil
	}
	return ErrRouterIsNotRunning
}

func (q Queue) HandlerRequest(msg *message.Message) error {
	fmt.Println(msg)
	return nil
}

func (q Queue) HandlerResponse(msg *message.Message) error {
	fmt.Println(msg)
	return nil
}

func main() {
	queue := Queue{
		mapRouter: make(map[string]*Router),
		locker:    sync.Mutex{},
	}
	msg := message.NewMessage(watermill.NewUUID(), []byte("Toi ten la dat"))
	err := queue.SendMessage("dat_in", msg)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	msg = message.NewMessage(watermill.NewUUID(), []byte("Ban co bi dien khong"))
	err = queue.SendMessage("dat_in", msg)
	if err != nil {
		panic(err)
	}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT)
	<-signalCh
	fmt.Println("Received Ctrl+C. Stopping the application gracefully.")
}
