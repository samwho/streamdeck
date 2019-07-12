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

type EventHandler func(ctx context.Context, client *Client, event EventReceived) error

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
			_, message, err := client.c.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}

			event := EventReceived{}
			if err := json.Unmarshal(message, &event); err != nil {
				log.Printf("failed to unmarshal received event: %s\n", string(message))
				return
			}

			ctx := setContext(ctx, event.Context)
			for _, f := range client.handlers[event.Event] {
				f(ctx, client, event)
			}

			log.Println("recv: ", string(message))
		}
	}()

	if err := client.register(params); err != nil {
		return nil, err
	}
	return client, nil
}

func (client *Client) register(params RegistrationParams) error {
	log.Println("sending register event...")
	if err := client.c.WriteJSON(NewRegisterEvent(params)); err != nil {
		client.Close()
		return err
	}
	return nil
}

func (client *Client) SetSettings(ctx context.Context, settings interface{}) error {
	return client.c.WriteJSON(NewEvent(ctx, SetSettings, settings))
}

func (client *Client) GetSettings(ctx context.Context) error {
	return client.c.WriteJSON(NewEvent(ctx, GetSettings, nil))
}

func (client *Client) Log(message string) error {
	return client.c.WriteJSON(NewLogMessage(message))
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
