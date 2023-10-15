package main

import (
	"context"
	"time"

	"github.com/sandronister/event/pkg/rabbimq"
)

func main() {
	ch := rabbimq.OpenChannel()
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rabbimq.Publish(ctx, ch, "Hello World!", "amq.direct")

}
