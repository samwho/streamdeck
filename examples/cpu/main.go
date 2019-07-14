package main

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/samwho/streamdeck"
	sdcontext "github.com/samwho/streamdeck/context"
	"github.com/shirou/gopsutil/cpu"
)

const (
	logFile = "C:\\Users\\samwh\\AppData\\Roaming\\Elgato\\StreamDeck\\logs\\cpu.log"

	imgX = 72
	imgY = 72
)

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
	setup(client)

	return client.Run()
}

func setup(client *streamdeck.Client) {
	action := client.Action("dev.samwho.streamdeck.cpu")

	contexts := make(map[string]struct{})

	action.RegisterHandler(streamdeck.WillAppear, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
		contexts[event.Context] = struct{}{}
		return nil
	})

	action.RegisterHandler(streamdeck.WillDisappear, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
		delete(contexts, event.Context)
		return nil
	})

	readings := make([]float64, imgX, imgX)

	go func() {
		for range time.Tick(time.Second) {
			for i := 0; i < imgX-1; i++ {
				readings[i] = readings[i+1]
			}

			r, err := cpu.Percent(0, false)
			if err != nil {
				log.Printf("error getting CPU reading: %v\n", err)
			}
			readings[imgX-1] = r[0]

			for ctxStr := range contexts {
				ctx := context.Background()
				ctx = sdcontext.WithContext(ctx, ctxStr)

				img, err := streamdeck.Image(graph(readings))
				if err != nil {
					log.Printf("error creating image: %v\n", err)
					continue
				}

				if err := client.SetImage(ctx, img, streamdeck.HardwareAndSoftware); err != nil {
					log.Printf("error setting image: %v\n", err)
					continue
				}

				if err := client.SetTitle(ctx, fmt.Sprintf("CPU\n%d%%", int(r[0])), streamdeck.HardwareAndSoftware); err != nil {
					log.Printf("error setting title: %v\n", err)
					continue
				}
			}
		}
	}()
}

func graph(readings []float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, imgX, imgY))
	for x := 0; x < imgX; x++ {
		reading := readings[x] / 100
		upto := int(float64(imgY) * reading)
		for y := 0; y < upto; y++ {
			img.Set(x, imgY-y, color.RGBA{R: 255, A: 255})
		}
		for y := upto; y < imgY; y++ {
			img.Set(x, imgY-y, color.Black)
		}
	}
	return img
}
