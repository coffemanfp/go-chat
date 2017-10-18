package domain

import (
	"testing"
)

func TestRoomCreated(t *testing.T) {
	r, ev := NewRoom("test", NewUserIDSet(1))
	// check whether room has one event,
	events := r.Events()
	if got := len(events); got != 1 {
		t.Errorf("room has no event after RoomCreated")
	}
	if _, ok := events[0].(RoomCreated); !ok {
		t.Errorf("invalid event state for the room")
	}

	// check whether room created event is valid.
	if got := ev.Name; got != "test" {
		t.Errorf("RoomCreated has different room name, expect: %s, got: %s", "test", got)
	}
	if got := len(ev.MemberIDs); got != 1 {
		t.Errorf("RoomCreated has dieffrent room members size, expect: %d, got: %d", 1, got)
	}

}

func TestRoomAddMember(t *testing.T) {
	r, _ := NewRoom("test", NewUserIDSet())
	r.ID = 1 // it may not be allowed at application side.
	u := User{ID: 1}
	ev, err := r.AddMember(u)
	if err != nil {
		t.Fatal(err)
	}
	if got := ev.RoomID; got != r.ID {
		t.Errorf("RoomAddedMember has different room id, expect: %d, got: %d", r.ID, got)
	}
	if got := ev.AddedUserID; got != u.ID {
		t.Errorf("RoomAddedMember has different user id, expect: %d, got: %d", u.ID, got)
	}

	if !r.HasMember(u) {
		t.Errorf("AddMember does not add any member to the room")
	}

	// room has two events: Created, AddedMember.
	if got := len(r.Events()); got != 2 {
		t.Errorf("room has no event")
	}
	if _, ok := r.Events()[1].(RoomAddedMember); !ok {
		t.Errorf("invalid event is added")
	}
}