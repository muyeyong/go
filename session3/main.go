package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
)

// manager 整个有一个过期时间

type Session struct {
	ID string
	Values    map[string]interface{}
	ExpiresAt time.Time
}

// 通过内存读取的provider
type Provider interface {
	SessionInit()(Session, error)
	SessionRead(sid string)(Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Manager struct {
	maxLifeTime int64
	// 类似于从哪里获取session, 内存或是数据库
	provider Provider
}

var provides = make(map[string]Provider)
func NewManager(providerName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	return &Manager{provider: provider, maxLifeTime: maxLifeTime}, nil
}

type memoryProvider struct {
	sync.Map
}


func GenerateSessionID() string {
	return uuid.Must(uuid.NewV4()).String()
}
func (m *memoryProvider) SessionInit()(*Session, error) {
	sessionId := GenerateSessionID()
	expiresAt := time.Now().Add(24 * time.Hour)
	session := &Session {
		ID: sessionId,
		Values: make(map[string]interface{}),
		ExpiresAt: expiresAt,
	}
	return session, nil
}

func (m *memoryProvider) SessionRead(sid string) (*Session, error) {
	//TODO ele是什么意思
	if ele, ok := m.Load(sid); ok {
		session, ok := ele.(*Session)
		if ok && session.ExpiresAt.After(time.Now()) {
			return session, nil
		}
	}
	return nil, fmt.Errorf("id `%s` not exist session data", sid)
}