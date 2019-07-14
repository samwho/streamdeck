package main

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"

	"github.com/samwho/streamdeck"
	"github.com/samwho/streamdeck/payload"
)

const (
	logFile = "C:\\Users\\samwh\\AppData\\Roaming\\Elgato\\StreamDeck\\logs\\streamdeck-livesplit.log"
)

type Settings struct {
	Counter int `json:"counter"`
}

func main() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func run(ctx context.Context) error {
	params, err := streamdeck.ParseRegistrationParams(os.Args)
	if err != nil {
		return err
	}

	client := streamdeck.NewClient(ctx, params)
	setupCounter(client)

	return client.Run()
}

func setupCounter(client *streamdeck.Client) {
	action := client.Action("dev.samwho.streamdeck.counter")
	settings := make(map[string]*Settings)

	action.RegisterHandler(streamdeck.WillAppear, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
		p := payload.WillAppear{}
		if err := json.Unmarshal(event.Payload, &p); err != nil {
			return err
		}

		s, ok := settings[event.Context]
		if !ok {
			s = &Settings{}
			settings[event.Context] = s
		}

		if err := json.Unmarshal(p.Settings, s); err != nil {
			return err
		}

		bg, err := streamdeck.Image(background())
		if err != nil {
			return err
		}

		if err := client.SetImage(ctx, bg, payload.HardwareAndSoftware); err != nil {
			return err
		}

		return client.SetTitle(ctx, strconv.Itoa(s.Counter), payload.HardwareAndSoftware)
	})

	action.RegisterHandler(streamdeck.WillDisappear, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
		s, _ := settings[event.Context]
		s.Counter = 0
		return client.SetSettings(ctx, s)
	})

	action.RegisterHandler(streamdeck.KeyDown, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
		s, ok := settings[event.Context]
		if !ok {
			return fmt.Errorf("couldn't find settings for context %v", event.Context)
		}

		s.Counter++
		if err := client.SetSettings(ctx, s); err != nil {
			return err
		}

		return client.SetTitle(ctx, strconv.Itoa(s.Counter), payload.HardwareAndSoftware)
	})
}

func background() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 20, 20))
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			img.Set(x, y, color.Black)
		}
	}
	return img
}
