package domain

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/shirasudon/go-chat/domain/event"
)

type UserRepositoryStub struct{}

func (u *UserRepositoryStub) BeginTx(context.Context, *sql.TxOptions) (Tx, error) {
	panic("not implemented")
}

func (uu *UserRepositoryStub) Store(ctx context.Context, u User) (uint64, error) {
	return u.ID + 1, nil
}

func (u *UserRepositoryStub) Find(ctx context.Context, id uint64) (User, error) {
	panic("not implemented")
}

var userRepo = &UserRepositoryStub{}

func TestUserCreated(t *testing.T) {
	ctx := context.Background()
	u, err := NewUser(ctx, userRepo, "user", "u-", "ser", "password", NewUserIDSet(1))
	if err != nil {
		t.Fatal(err)
	}

	if u.ID == 0 {
		t.Fatalf("user is created but has invalid ID(%d)", u.ID)
	}

	// check whether user has one event,
	events := u.Events()
	if got := len(events); got != 1 {
		t.Errorf("user has no event after UserCreated")
	}
	ev, ok := events[0].(event.UserCreated)
	if !ok {
		t.Errorf("invalid event state for the user")
	}

	// check whether user created event is valid.
	if got := ev.Name; got != "user" {
		t.Errorf("UserCreated has different user name, expect: %s, got: %s", "user", got)
	}
	if got := ev.FirstName; got != "u-" {
		t.Errorf("UserCreated has different first name, expect: %s, got: %s", "u-", got)
	}
	if got := ev.LastName; got != "ser" {
		t.Errorf("UserCreated has different last name, expect: %s, got: %s", "ser", got)
	}
	if got := len(ev.FriendIDs); got != 1 {
		t.Errorf("UseerCreated has dieffrent friends size, expect: %d, got: %d", 1, got)
	}
	if got := ev.Timestamp(); got == (time.Time{}) {
		t.Error("UserCreated has no timestamp")
	}

}

func TestUserAddFriendSuccess(t *testing.T) {
	ctx := context.Background()
	u, _ := NewUser(ctx, userRepo, "user", "u-", "ser", "password", NewUserIDSet())
	u.ID = 1 // it may not be allowed at application side.
	friend := User{ID: u.ID + 1}
	ev, err := u.AddFriend(friend)
	if err != nil {
		t.Fatal(err)
	}
	if got := ev.UserID; got != u.ID {
		t.Errorf("UserAddedFriend has different user id, expect: %d, got: %d", u.ID, got)
	}
	if got := ev.AddedFriendID; got != friend.ID {
		t.Errorf("UserAddedFriend has different friend id, expect: %d, got: %d", friend.ID, got)
	}
	if got := ev.Timestamp(); got == (time.Time{}) {
		t.Error("UserAddedFriend has no timestamp")
	}

	if !u.HasFriend(friend) {
		t.Errorf("AddFriend could not add friend to the user")
	}

	// user has two events: Created, AddedFriend.
	if got := len(u.Events()); got != 2 {
		t.Errorf("user has no event")
	}
	if _, ok := u.Events()[1].(event.UserAddedFriend); !ok {
		t.Errorf("invalid event is added")
	}
}

func TestUserAddFriendFail(t *testing.T) {
	// fail case: Add itself as friend.
	ctx := context.Background()
	u, _ := NewUser(ctx, userRepo, "user", "u-", "ser", "password", NewUserIDSet())
	u.ID = 1 // it may not be allowed at application side.
	_, err := u.AddFriend(u)
	if err == nil {
		t.Fatal("add itself as friend but no error")
	}

	// user has one events: Created.
	if got := len(u.Events()); got != 1 {
		t.Errorf("user has invalid event state")
	}
}
