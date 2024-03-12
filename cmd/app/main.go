package main

import (
	"context"
	"log"

	"github.com/dhucsik/bookers/internal/app"
)

func main() {
	ctx := context.Background()
	a := app.InitApp(ctx)

	if err := a.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
