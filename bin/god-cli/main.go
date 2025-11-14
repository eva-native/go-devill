package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	godevill "github.com/eva-native/go-devill"
)

var token = flag.String("token", "", "godville API token")

func main() {
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	godville, err := godevill.New("Tungrad", godevill.WithToken(*token))
	if err != nil {
		log.Fatalln(err)
	}

	err = poll(ctx, godville)
	if err != nil {
		log.Fatalln(err)
	}
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := poll(ctx, godville); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func poll(ctx context.Context, api *godevill.API) error {
	timeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	hero, err := api.GetWithContext(timeout)
	if err != nil {
		return err
	}
	text, err := json.MarshalIndent(hero, "", "  ")
	if err != nil {
		return err
	}
	log.Println(string(text))
	return nil
}
