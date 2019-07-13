package streamdeck

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type EventHandler func(ctx context.Context, client *Client, event Event) error

type Client struct {
	c        *websocket.Conn
	handlers map[string][]EventHandler
	done     chan struct{}
}

func NewClient(ctx context.Context, params RegistrationParams) (*Client, error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("127.0.0.1:%d", params.Port)}
	log.Printf("connecting to StreamDeck at %v\n", u)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	log.Printf("connected to StreamDeck\n")

	done := make(chan struct{})

	client := &Client{
		c:        c,
		handlers: make(map[string][]EventHandler),
		done:     done,
	}

	go func() {
		defer close(done)
		log.Println("starting read loop")
		for {
			messageType, message, err := client.c.ReadMessage()
			if err != nil {
				log.Printf("read error: %v\n", err)
				return
			}

			if messageType == websocket.PingMessage {
				log.Printf("received ping message\n")
				if err := client.c.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
					log.Printf("error while ponging: %v\n", err)
				}
				continue
			}

			if messageType == websocket.CloseMessage {
				// handle close message
				panic("websocket close!")
			}

			event := Event{}
			if err := json.Unmarshal(message, &event); err != nil {
				log.Printf("failed to unmarshal received event: %s\n", string(message))
				continue
			}

			log.Println("recv: ", string(message))

			ctx := setContext(ctx, event.Context)
			for _, f := range client.handlers[event.Event] {
				if err := f(ctx, client, event); err != nil {
					log.Printf("error in handler for event %v: %v\n", event.Event, err)
				}
			}
		}
	}()

	if err := client.register(params); err != nil {
		return nil, err
	}
	return client, nil
}

func (client *Client) register(params RegistrationParams) error {
	log.Println("sending register event...")
	if err := client.send(Event{UUID: params.PluginUUID, Event: params.RegisterEvent}); err != nil {
		client.Close()
		return err
	}
	return nil
}

func (client *Client) send(event Event) error {
	j, _ := json.Marshal(event)
	log.Printf("sending message: %v\n", string(j))
	return client.c.WriteJSON(event)
}

func (client *Client) SetSettings(ctx context.Context, settings interface{}) error {
	return client.send(NewEvent(ctx, SetSettings, settings))
}

func (client *Client) GetSettings(ctx context.Context) error {
	return client.send(NewEvent(ctx, GetSettings, nil))
}

func (client *Client) SetGlobalSettings(ctx context.Context, settings interface{}) error {
	return client.send(NewEvent(ctx, SetGlobalSettings, settings))
}

func (client *Client) GetGlobalSettings(ctx context.Context) error {
	return client.send(NewEvent(ctx, GetGlobalSettings, nil))
}

func (client *Client) OpenURL(ctx context.Context, u url.URL) error {
	return client.send(NewEvent(ctx, OpenURL, OpenURLPayload{URL: u.String()}))
}

func (client *Client) LogMessage(message string) error {
	return client.send(NewEvent(nil, LogMessage, LogMessagePayload{Message: message}))
}

func (client *Client) SetTitle(ctx context.Context, title string, target Target) error {
	return client.send(NewEvent(ctx, SetTitle, SetTitlePayload{Title: title, Target: target}))
}

func (client *Client) SetImage(ctx context.Context, base64image string, target Target) error {
	return client.send(NewEvent(ctx, SetImage, SetImagePayload{Base64Image: base64image, Target: target}))
}

func (client *Client) RegisterHandler(eventName string, handler EventHandler) {
	client.handlers[eventName] = append(client.handlers[eventName], handler)
}

func (client *Client) Close() error {
	err := client.c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	select {
	case <-client.done:
	case <-time.After(time.Second):
	}
	return client.c.Close()
}

func (client *Client) Join() {
	<-client.done
}
