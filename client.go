package imsgbridge

import (
	"context"
	"strings"
	"sync"

	"github.com/ninepeach/imsgbridge/internal/applemsg"
	"github.com/ninepeach/imsgbridge/internal/sender"
	"github.com/ninepeach/imsgbridge/internal/state"
)

type MessageHandler func(Message)

type Client struct {
	service *applemsg.Service

	sender sender.Sender

	store *state.Store

	mu sync.RWMutex

	handlers []MessageHandler

	assistantPrefix string

	systemPrefix string

	notificationPrefix string
}

func NewClient(
	config Config,
) (*Client, error) {

	if config.AssistantPrefix == "" {

		config.AssistantPrefix = "[assistant]"

	}

	if config.SystemPrefix == "" {

		config.SystemPrefix = "[system]"

	}

	if config.NotificationPrefix == "" {

		config.NotificationPrefix = "[notification]"

	}

	db, err := applemsg.OpenDB()

	if err != nil {

		return nil, err

	}

	reader := applemsg.NewReader(
		db,
	)

	service := applemsg.NewService(
		reader,
	)

	store, err := state.NewStore()

	if err != nil {

		return nil, err

	}

	return &Client{

		service: service,

		sender: sender.New(),

		store: store,

		handlers: make(
			[]MessageHandler,
			0,
		),

		assistantPrefix: config.AssistantPrefix,

		systemPrefix: config.SystemPrefix,

		notificationPrefix: config.NotificationPrefix,
	}, nil

}

// OnMessage registers a message handler.
func (c *Client) OnMessage(
	handler MessageHandler,
) {

	c.mu.Lock()

	defer c.mu.Unlock()

	c.handlers = append(
		c.handlers,
		handler,
	)

}

// Start starts receiving messages.
func (c *Client) Start() error {


	ctx := context.Background()


	return c.service.Run(
		ctx,
		c.store,

		func(msg applemsg.Message) {


			// ignore empty messages
			if msg.Text == "" {

				return

			}


			// ignore assistant/system generated messages
			if c.isInternalMessage(
				msg.Text,
			) {

				return

			}



			publicMessage := Message{

				ID: msg.ID,

				From: msg.Sender.Handle,

				Text: msg.Text,

				Time: msg.Time,

				IsFromMe:
					msg.Direction ==
					applemsg.Outgoing,

			}


			c.dispatch(
				publicMessage,
			)

		},
	)

}

func (c *Client) dispatch(
	msg Message,
) {

	c.mu.RLock()

	defer c.mu.RUnlock()

	for _, handler := range c.handlers {

		go handler(msg)

	}

}

// Send sends a normal message.
func (c *Client) Send(
	target string,
	text string,
) error {

	return c.sender.Send(
		target,
		text,
	)

}

// SendAssistant sends an assistant response.
func (c *Client) SendAssistant(
	target string,
	text string,
) error {

	return c.sendTagged(
		target,
		c.assistantPrefix,
		text,
	)

}

// SendSystem sends a system message.
func (c *Client) SendSystem(
	target string,
	text string,
) error {

	return c.sendTagged(
		target,
		c.systemPrefix,
		text,
	)

}

// SendNotification sends a notification message.
func (c *Client) SendNotification(
	target string,
	text string,
) error {

	return c.sendTagged(
		target,
		c.notificationPrefix,
		text,
	)

}

func (c *Client) sendTagged(
	target string,
	prefix string,
	text string,
) error {

	if prefix != "" {

		text =
			prefix +
				" " +
				text

	}

	return c.sender.Send(
		target,
		text,
	)

}

func (c *Client) isInternalMessage(
	text string,
) bool {

	return strings.HasPrefix(
		text,
		c.assistantPrefix,
	) ||

		strings.HasPrefix(
			text,
			c.systemPrefix,
		) ||

		strings.HasPrefix(
			text,
			c.notificationPrefix,
		)

}
