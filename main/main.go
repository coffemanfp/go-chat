// +build !appengine

package main

import (
	"context"
	"log"

	gochat "github.com/shirasudon/go-chat"
	"github.com/shirasudon/go-chat/chat"
	"github.com/shirasudon/go-chat/infra/inmemory"
	"github.com/shirasudon/go-chat/infra/pubsub"
)

func createServer() (server *gochat.Server, done func()) {
	ps := pubsub.New()
	// defer ps.Shutdown()

	repos := inmemory.OpenRepositories(ps)
	// defer repos.Close()

	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	go repos.UpdatingService(ctx)

	qs := &chat.Queryers{
		UserQueryer:    repos.UserRepository,
		RoomQueryer:    repos.RoomRepository,
		MessageQueryer: repos.MessageRepository,
		EventQueryer:   repos.EventRepository,
	}

	server = gochat.NewServer(repos, qs, ps, nil)
	done = func() {
		cancel()
		repos.Close()
		ps.Shutdown()
	}
	return
}

func main() {
	s, done := createServer()
	defer done()
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
