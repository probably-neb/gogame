package sessions

import (
    "crypto/rand"
    "io"
    "encoding/base64"
    "sync"
    "errors"
)

type Manager struct {
    lock sync.RWMutex
	sessions map[string]Session
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]Session),
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
    m.sessions[id] = Session{Exists: true}
	return id
}

func (m *Manager) Get(id string) Session {
    m.lock.RLock()
    defer m.lock.RUnlock()
    return m.sessions[id]
}

func (m *Manager) Set(id string, field string, value any) error {
    m.lock.Lock()
    defer m.lock.Unlock()
    if field == "Name" {
        value, ok := value.(string)
        if !ok {
            return errors.New("tried to set session var name to non string value")
        }
        session := m.sessions[id]
        if !session.Exists {
            return errors.New("session with sessionId="+ id+ " does not exist")
        }
        session.Name = value
        return nil
    }
    return errors.New("tried to set unknown field: "+ field+ " in sessions")
}

type Session struct {
    // will be false
    Exists bool
	Name string
}
