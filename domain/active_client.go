package domain

import (
	"errors"
	"fmt"
	"sync"
)

// ActiveClientRepository manages active clients.
// It is inmemory struct type rather than interface
// because ActiveClient has connection interface
// which can not be serialized.
type ActiveClientRepository struct {
	clientsMu *sync.RWMutex
	clients   map[uint64]*ActiveClient
}

func NewActiveClientRepository(capHint int) *ActiveClientRepository {
	return &ActiveClientRepository{
		clientsMu: new(sync.RWMutex),
		clients:   make(map[uint64]*ActiveClient, capHint),
	}
}

// Find all ActiveClients by user ID list.
// It returns found AcitiveClients and error if all of the user
// ID are not found.
// The number of the returned ActiveClients may not same as
// that of user ID list, when a part of the user ID list are not
// found.
func (cm *ActiveClientRepository) FindAllByUserIDs(userIDs []uint64) ([]*ActiveClient, error) {
	friends := make([]*ActiveClient, 0, 8)

	cm.clientsMu.RLock()
	defer cm.clientsMu.RUnlock()

	for _, id := range userIDs {
		if activeC, ok := cm.clients[id]; ok {
			friends = append(friends, activeC)
		}
	}

	if len(friends) == 0 {
		return nil, fmt.Errorf("no found any ActiveClients (%v)", userIDs)
	}
	return friends, nil
}

// Find ActiveClient by user ID.
// It returns found AcitiveClient and error if not found.
func (cm *ActiveClientRepository) Find(userID uint64) (*ActiveClient, error) {
	cm.clientsMu.RLock()
	defer cm.clientsMu.RUnlock()

	if activeC, ok := cm.clients[userID]; ok {
		return activeC, nil
	}
	return nil, fmt.Errorf("active user(%d) is not found", userID)
}

// Check whether the specified connection exists in the
// repository.
// It returns true if found otherwise returns false.
func (cm *ActiveClientRepository) ExistByConn(c Conn) bool {
	cm.clientsMu.RLock()
	activeC, ok := cm.clients[c.UserID()]
	cm.clientsMu.RUnlock()

	if ok {
		return activeC.HasConn(c)
	}
	return false
}

// Store activeClient to the repository. It returns error
// when something wrong.
func (cm *ActiveClientRepository) Store(ac *ActiveClient) error {
	cm.clientsMu.Lock()
	defer cm.clientsMu.Unlock()

	// just update connection list if ActiveClient already exist.
	if activeC, ok := cm.clients[ac.userID]; ok {
		activeC.conns = ac.conns
		return nil
	}

	// store new active client because the ActiveClient is newly.
	cm.clients[ac.userID] = ac
	return nil
}

// Remove connection from the repository.
func (cm *ActiveClientRepository) Remove(ac *ActiveClient) error {
	if ac == nil {
		return errors.New("can not remove nil ActiveClient")
	}

	cm.clientsMu.Lock()
	defer cm.clientsMu.Unlock()

	_, ok := cm.clients[ac.userID]
	if !ok {
		return fmt.Errorf("active user(%d) is not found", ac.userID)
	}

	delete(cm.clients, ac.userID)
	return nil
}

// connection specific infomation.
type connectionInfo struct {
	// TODO
}

// ActiveClient is a one active domain User which has several
// connections.
//
// ActiveClient can have multiple conncetions,
// because one user can conncets from PC,
// mobile device, and so on.
type ActiveClient struct {
	userID uint64

	mu    *sync.RWMutex
	conns map[Conn]connectionInfo
}

func NewActiveClient(repo *ActiveClientRepository, c Conn, u User) (*ActiveClient, ActiveClientActivated, error) {
	if u.ID != c.UserID() {
		return nil, ActiveClientActivated{}, fmt.Errorf("can not create AcitiveClient with different user(%d) and conncetion(userID=%d)", u.ID, c.UserID())
	}
	if _, err := repo.Find(c.UserID()); err == nil {
		return nil, ActiveClientActivated{}, fmt.Errorf("user(%d) is already active", c.UserID())
	}

	ac := &ActiveClient{
		userID: c.UserID(),
		mu:     new(sync.RWMutex),
		conns:  map[Conn]connectionInfo{c: connectionInfo{}},
	}
	err := repo.Store(ac)
	if err != nil {
		return nil, ActiveClientActivated{}, err
	}

	ev := ActiveClientActivated{
		UserID:   c.UserID(),
		UserName: "", // TODO
	}
	return ac, ev, nil
}

// Delete deletes this ActiveClient from the repository.
func (ac *ActiveClient) Delete(repo *ActiveClientRepository) (ActiveClientInactivated, error) {
	ac.mu.RLock()
	if len(ac.conns) > 0 {
		ac.mu.RUnlock()
		return ActiveClientInactivated{}, errors.New("AcitiveClient contains any connection, can not be deleted")
	}
	ac.mu.RUnlock()

	err := repo.Remove(ac)
	if err != nil {
		return ActiveClientInactivated{}, fmt.Errorf("AcitiveClient not in the repository, can not be deleted: %v", err)
	}

	return ActiveClientInactivated{
		UserID:   ac.userID,
		UserName: "", // TODO
	}, nil
}

func (ac *ActiveClient) HasConn(c Conn) bool {
	ac.mu.RLock()
	_, exist := ac.conns[c]
	ac.mu.RUnlock()
	return exist
}

// Send domain event to all of the client connections.
func (ac *ActiveClient) Send(ev Event) {
	ac.mu.RLock()
	for c, _ := range ac.conns {
		c.Send(ev)
	}
	ac.mu.RUnlock()
}

// MaxConns is the maximum number of the connections
// per ActiveClient.
const MaxConns = 16

// ErrExceedsConnsMax indicates the number of the connection for the
// ActiveClient exceeds the maximum limits.
var ErrExceedsMaxConns = errors.New("exceed the number of connections")

// Add new connection to ActiveClient.
// It returns current number of the connections and
// error when the connection is not for the AcitiveClient or
// ErrExceedsMaxConns when the number of the connections exceeds
// the maximum limit.
func (ac *ActiveClient) AddConn(c Conn) (int, error) {
	if ac.userID != c.UserID() {
		return 0, fmt.Errorf("AcitiveClient(%d) can not contain the different user connection (userID=%d)", ac.userID, c.UserID())
	}

	ac.mu.Lock()
	defer ac.mu.Unlock()

	if len(ac.conns) == MaxConns {
		return MaxConns, ErrExceedsMaxConns
	}

	ac.conns[c] = connectionInfo{}
	return len(ac.conns), nil
}

// RemoveConn removes connection from ActiveClient.
// It returns the rest number of the connection and
// errors related with the connection is removed.
func (ac *ActiveClient) RemoveConn(c Conn) (int, error) {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	if _, ok := ac.conns[c]; !ok {
		return len(ac.conns), errors.New("connection is not found")
	}

	delete(ac.conns, c)
	return len(ac.conns), nil
}

// domain event for the AcitiveClient is activated.
type ActiveClientActivated struct {
	UserID   uint64
	UserName string
}

func (ActiveClientActivated) EventType() EventType { return EventActiveClientActivated }

// domain event for the AcitiveClient is inactivated.
type ActiveClientInactivated struct {
	UserID   uint64
	UserName string
}

func (ActiveClientInactivated) EventType() EventType { return EventActiveClientInactivated }
