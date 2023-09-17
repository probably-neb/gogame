package sessions

import (
    "crypto/rand"
    "io"
    "encoding/base64"
    "sync"
    "errors"
	"gogame/random-name"
)

type Manager struct {
    lock sync.RWMutex
	sessions map[string]*Session
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]*Session),
	}
}

func (m *Manager) NewSession() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
    id := base64.URLEncoding.EncodeToString(b)
    m.lock.Lock()
    defer m.lock.Unlock()
	name := randomName.RandomName()
	m.sessions[id] = &Session{Name: name}
	return id
}

func (m *Manager) Get(id string) (Session, bool) {
    m.lock.RLock()
    defer m.lock.RUnlock()
    session, ok := m.sessions[id]
    if ok {
        return *session, ok
    }
    return Session{}, ok
}

func (m *Manager) Set(id string, field string, value any) error {
    m.lock.Lock()
    defer m.lock.Unlock()
    if field == "Name" {
        value, ok := value.(string)
        if !ok {
            return errors.New("tried to set session var name to non string value")
        }
        session, ok := m.sessions[id]
        if !ok {
            return errors.New("session with sessionId="+ id+ " does not exist")
        }
        session.Name = value
        return nil
    }
    return errors.New("tried to set unknown field: "+ field+ " in sessions")
}

type Session struct {
	Name string
}
