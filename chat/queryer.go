package chat

import (
	"context"
	"time"

	"github.com/shirasudon/go-chat/chat/queried"
	"github.com/shirasudon/go-chat/domain"
	"github.com/shirasudon/go-chat/domain/event"
)

//go:generate mockgen -destination=../internal/mocks/mock_queryer.go -package=mocks github.com/shirasudon/go-chat/chat UserQueryer,RoomQueryer,MessageQueryer,EventQueryer

// Queryers is just data struct which have
// some XXXQueryers.
type Queryers struct {
	UserQueryer
	RoomQueryer
	MessageQueryer

	EventQueryer
}

// UserQueryer queries users stored in the data-store.
type UserQueryer interface {
	// Find a user specified by userID and return it.
	// It returns NotFoundError if not found.
	Find(ctx context.Context, userID uint64) (domain.User, error)

	// Find a user profile specified by user name and password.
	// It returns error if not found.
	FindByNameAndPassword(ctx context.Context, name, password string) (*queried.AuthUser, error)

	// Find a user related information with userID.
	// It returns queried result or NotFoundError if the information is not found.
	FindUserRelation(ctx context.Context, userID uint64) (*queried.UserRelation, error)
}

// RoomQueryer queries rooms stored in the data-store.
type RoomQueryer interface {
	// Find a room specified by roomID and return it.
	// It returns NotFoundError if not found.
	Find(ctx context.Context, roomID uint64) (domain.Room, error)

	// Find all rooms which user has.
	// It returns NotFoundError if not found.
	FindAllByUserID(ctx context.Context, userID uint64) ([]domain.Room, error)

	// Find room information with specified userID and roomID.
	// It returns NotFoundError if not found.
	FindRoomInfo(ctx context.Context, userID, roomID uint64) (*queried.RoomInfo, error)
}

// MessageQueryer queries messages stored in the data-store.
type MessageQueryer interface {

	// Find all messages from the room specified by room_id.
	// The returned messages are, ordered by latest created at,
	// all of before specified before time,
	// and the number of messages is limted to less than
	// specified limit.
	// It returns InfraError if infrastructure raise some errors.
	// It returns NotFoundError if not found.
	FindRoomMessagesOrderByLatest(ctx context.Context, roomID uint64, before time.Time, limit int) ([]domain.Message, error)

	// Find all unread messages from the room specified by room_id.
	// The returned messages are, ordered by latest created at,
	// It returns NotFoundError if not found.
	FindUnreadRoomMessages(ctx context.Context, userID, roomID uint64, limit int) (*queried.UnreadRoomMessages, error)
}

// EventQueryer queries events stored in the data-store.
type EventQueryer interface {
	// Find events from the data-store.
	// The returned events are, ordered by older created at
	// and all of after specified after time.
	// It returns NotFoundError if not found.
	FindAllByTimeCursor(ctx context.Context, after time.Time, limit int) ([]event.Event, error)

	// Find events, associated with specified stream ID, from the data-store.
	// The returned events are, ordered by older created at
	// and all of after specified after time.
	// It returns NotFoundError if not found.
	FindAllByStreamID(ctx context.Context, streamID event.StreamID, after time.Time, limit int) ([]event.Event, error)
}
