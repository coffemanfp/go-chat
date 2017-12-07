package main

import (
	"context"
	"log"

	gochat "github.com/shirasudon/go-chat"
	"github.com/shirasudon/go-chat/chat"
	"github.com/shirasudon/go-chat/infra/inmemory"
	"github.com/shirasudon/go-chat/infra/pubsub"
)

func main() {
	// initilize database

	ps := pubsub.New()
	defer ps.Shutdown()

	repos := inmemory.OpenRepositories(ps)
	defer repos.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go repos.UpdatingService(ctx)

	qs := &chat.Queryers{
		UserQueryer:    repos.UserRepository,
		RoomQueryer:    repos.RoomRepository,
		MessageQueryer: repos.MessageRepository,
		EventQueryer:   repos.EventRepository,
	}
	log.Fatal(gochat.ListenAndServe(repos, qs, ps, nil))
}
