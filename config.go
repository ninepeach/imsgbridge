package imsgbridge

import "time"

type Config struct {

	// PollInterval controls message polling frequency.
	PollInterval time.Duration

	// Prefix for assistant generated messages.
	AssistantPrefix string

	// Prefix for internal system messages.
	SystemPrefix string

	// Prefix for notification messages.
	NotificationPrefix string
}
