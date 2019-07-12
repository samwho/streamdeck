package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/samwho/streamdeck-livesplit/streamdeck"
)

const (
	logFile = "C:\\Users\\samwh\\AppData\\Roaming\\Elgato\\StreamDeck\\logs\\streamdeck-livesplit.log"
)

var (
	port          = flag.Int("port", -1, "")
	pluginUUID    = flag.String("pluginUUID", "", "")
	registerEvent = flag.String("registerEvent", "", "")
	info          = flag.String("info", "", "")
)

func main() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	flag.Parse()

	params := streamdeck.RegistrationParams{
		Port:          *port,
		PluginUUID:    *pluginUUID,
		RegisterEvent: *registerEvent,
		Info:          *info,
	}

	log.Printf("registration params: %v\n", params)
	ctx := context.Background()
	if err := run(ctx, params); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func run(ctx context.Context, params streamdeck.RegistrationParams) error {
	client, err := streamdeck.NewClient(ctx, params)
	if err != nil {
		return err
	}

	client.RegisterHandler(streamdeck.KeyDown, func(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) {
		return client.Log("key down!")
	})

	log.Println("waiting for connection to close...")
	client.Join()

	return nil
}
