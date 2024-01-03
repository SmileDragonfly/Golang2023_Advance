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
	"reflect"
	"sync"
	"syscall"
	"time"
)

var ErrRouterIsNotRunning error = errors.New("Router is not running")

type Queue struct {
	router *message.Router
	pub    message.Publisher
	sub    message.Subscriber
	locker sync.Mutex
}

func (q *Queue) AddRouter(name string) error {
	q.locker.Lock()
	defer q.locker.Unlock()
	if q.router != nil {
		return nil
	} else {
		router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
		if err != nil {
			return err
		}
		pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
		q.router = router
		q.sub = pubSub
		q.pub = pubSub
	}
	q.router.AddNoPublisherHandler(name, name, q.sub, q.HandlerRequest)
	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	ctx := context.Background()
	go q.router.Run(ctx)
	<-q.router.Running()
	return nil
}

func (q *Queue) SendMessage(name string, msg *message.Message) error {
	if q.router == nil {
		err := q.AddRouter(name)
		if err != nil {
			return err
		}
	}
	if q.router.IsRunning() {
		err := q.pub.Publish(name, msg)
		if err != nil {
			return err
		}
		return nil
	}
	return ErrRouterIsNotRunning
}

func (q *Queue) HandlerRequest(msg *message.Message) error {
	fmt.Println(msg)
	return nil
}

func (q *Queue) HandlerResponse(msg *message.Message) error {
	fmt.Println(msg)
	return nil
}

func main() {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
	go func() {
		for {
			msgCh, err := pubSub.Subscribe(context.Background(), "datin")
			if err != nil {
				panic(err)
			}
			msg := <-msgCh
			fmt.Println(reflect.TypeOf(msg))
			fmt.Println(string(msg.Payload))
		}
	}()
	time.Sleep(time.Second)
	err := pubSub.Publish("datin", &message.Message{
		UUID:     watermill.NewUUID(),
		Metadata: nil,
		Payload:  []byte("Toi ten la dat"),
	})
	time.Sleep(10 * time.Second)
	if err != nil {
		panic(err)
	}
	err = pubSub.Publish("datin", &message.Message{
		UUID:     watermill.NewUUID(),
		Metadata: nil,
		Payload:  []byte("abcd"),
	})
	if err != nil {
		panic(err)
	}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT)
	<-signalCh
}
