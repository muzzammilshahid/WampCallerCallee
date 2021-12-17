package main

import (
	"context"
	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
	"log"
	"os"
	"time"
)

func caller() {
	logger := log.New(os.Stderr, "CALLER> ", 0)

	cfg := client.Config{
		Realm:  "realm1",
		Logger: logger,
	}
	caller, err := client.ConnectNet(context.Background(), "ws://192.168.100.2:8080/ws", cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer caller.Close()

	callArgs := wamp.List{"Test"}
	ctx := context.Background()
	start := time.Now()

	for i := 1; i < 1000; i++ {
		result, err := caller.Call(ctx, procedureName, nil, callArgs, nil, nil)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Println(result.Arguments)
	}

	elapsed := time.Since(start)
	logger.Println("loop take ", elapsed)
}
