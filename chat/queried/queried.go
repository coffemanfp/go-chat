// package queried contains queried results as
// Data-Transfer-Object.

package queried

import "time"

// EmptyRoomInfo is RoomInfo having empty fields rather than nil.
var EmptyRoomInfo = RoomInfo{
	Members: []RoomMemberProfile{},
}

// RoomInfo is a detailed room information.
// creator
type RoomInfo struct {
	RoomName    string              `json:"room_name"`
	RoomID      uint64              `json:"room_id"`
	CreatorID   uint64              `json:"room_creator_id"`
	Members     []RoomMemberProfile `json:"room_members"`
	MembersSize int                 `json:"room_members_size"`
}

// RoomMemberProfile is a user profile with room specific information.
type RoomMemberProfile struct {
	UserProfile

	MessageReadAt time.Time `json:"message_read_at"`
}

// EmptyUserRelation is UserRelation having empty fields rather than nil.
var EmptyUserRelation = UserRelation{
	Friends: []UserProfile{},
	Rooms:   []UserRoom{},
}

// UserRelation is the abstarct information associated with specified User.
type UserRelation struct {
	UserProfile

	Friends []UserProfile `json:"friends"`
	Rooms   []UserRoom    `json:"rooms"`
}

// AuthUser is a authenticated user information.
type AuthUser struct {
	ID       uint64 `json:"user_id"`
	Name     string `json:"user_name"`
	Password string `json:"password"`
}

// UserProfile holds information for user profile.
type UserProfile struct {
	UserID    uint64 `json:"user_id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UserRoom holds abstract information for the room.
type UserRoom struct {
	RoomID   uint64 `json:"room_id"`
	RoomName string `json:"room_name"`
}

// EmptyRoomMessages is RoomMessages having empty fields rather than nil.
var EmptyRoomMessages = RoomMessages{
	Msgs: []Message{},
}

// RoomMessages is a message list in specified Room.
type RoomMessages struct {
	RoomID uint64 `json:"room_id"`

	Msgs []Message `json:"messages"`

	Cursor struct {
		Current time.Time `json:"current"`
		Next    time.Time `json:"next"`
	} `json:"cursor"`
}

type Message struct {
	MessageID uint64    `json:"message_id"`
	UserID    uint64    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// EmptyUnreadRoomMessages is UnreadRoomMessages having empty fields rather than nil.
var EmptyUnreadRoomMessages = UnreadRoomMessages{
	Msgs: []Message{},
}

// UnreadRoomMessages is a list of unread messages in specified Room.
type UnreadRoomMessages struct {
	RoomID uint64 `json:"room_id"`

	Msgs     []Message `json:"messages"`
	MsgsSize int       `json:"messages_size"`
}
