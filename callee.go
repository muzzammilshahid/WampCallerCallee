package main

import (
	"context"
	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/transport/serialize"
	"github.com/gammazero/nexus/v3/wamp"
	"log"
	"os"
	"os/signal"
)

func callee(url string, proc string) {
	logger := log.New(os.Stdout, "CALLEE> ", 0)
	cfg := client.Config{
		Realm:         "realm1",
		Logger:        logger,
		Serialization: serialize.CBOR,
	}
	// Connect callee client with requested socket type and serialization.
	callee, err := client.ConnectNet(context.Background(), url, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer callee.Close()

	// Register procedure
	if err = callee.Register(proc, result, nil); err != nil {
		logger.Fatal("Failed to register procedure:", err)
	}
	logger.Println("Registered procedure", proc, "with router")

	// Wait for CTRL-c or client close while handling remote procedure calls.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case <-sigChan:
	case <-callee.Done():
		logger.Print("Router gone, exiting")
		return // router gone, just exit
	}

	if err = callee.Unregister(procedureName); err != nil {
		logger.Println("Failed to unregister procedure:", err)
	}
}

func result(ctx context.Context, inv *wamp.Invocation) client.InvokeResult {
	log.Println("event receive")

	return client.InvokeResult{Args: wamp.List{}}
}
