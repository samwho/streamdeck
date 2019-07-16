package main

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/samwho/streamdeck"
)

type Settings struct {
	Counter int `json:"counter"`
}

func main() {
	f, err := ioutil.TempFile("", "streamdeck-counter.log")
	if err != nil {
		log.Fatalf("error creating temp file: %v", err)
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
	setup(client)

	return client.Run()
}

func setup(client *streamdeck.Client) {
	action := client.Action("dev.samwho.streamdeck.counter")
	settings := make(map[string]*Settings)

	action.RegisterHandler(streamdeck.WillAppear, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
		p := streamdeck.WillAppearPayload{}
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

		if err := client.SetImage(ctx, bg, streamdeck.HardwareAndSoftware); err != nil {
			return err
		}

		return client.SetTitle(ctx, strconv.Itoa(s.Counter), streamdeck.HardwareAndSoftware)
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

		return client.SetTitle(ctx, strconv.Itoa(s.Counter), streamdeck.HardwareAndSoftware)
	})
}

func background() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 72, 72))
	for x := 0; x < 72; x++ {
		for y := 0; y < 72; y++ {
			img.Set(x, y, color.Black)
		}
	}
	return img
}
