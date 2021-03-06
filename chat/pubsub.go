package chat

import (
	"github.com/shirasudon/go-chat/domain/event"
)

//go:generate mockgen -destination=../internal/mocks/mock_pubsub.go -package=mocks github.com/shirasudon/go-chat/chat Pubsub

// interface for the publisher/subcriber pattern.
type Pubsub interface {
	Pub(...event.Event)
	Sub(...event.Type) chan interface{}
}
