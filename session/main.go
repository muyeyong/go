package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(ley interface{}) error
	SessionID() string
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Manager struct {
	cookieName string
	lock       sync.Mutex
	provider  Provider
	maxLifeTime int64
}

type MemoryStore struct {
	sync.Map
}

type CSession struct {
	ID string
	Expires time.Time
	Date map[interface {}]interface{}
}

func (s *CSession) Get(key interface{}) (interface{}, error) {
	sess, err := globalSessions.provider.SessionRead(s.ID)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

var provides = make(map[string]Provider)
var globalSessions *Manager
func NewManager(provideName, cookieName string, maxLifeTime int64)(*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

func NewSession()(*CSession, error) {
	
}

func initSession() {
	globalSessions, _ = NewManager("memory", "sessionId", 3600)
}

func (m *MemoryStore)SessionInit(sid string) error {
	// 初始化session，将sid和session对应起来
	// 需要去new一个session
	// m.Store(sid, )
	return nil
}

func initProvide() {

}


func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _ , err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r * http.Request)(session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{ Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("session/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func main(){
	initSession()
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



