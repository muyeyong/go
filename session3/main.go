package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"runtime"
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
	SessionInit(sid string)(*Session, error)
	SessionRead(sid string)(*Session, error)
	SessionUpdate(sid string, value interface {})(error)
	SessionDestroy(sid string) error
	SessionGC()
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

func NewProvider(providerName string) {
	switch providerName {
	case "memory":
		provides["memory"] = new(memoryProvider)
	default:
		panic("session: unknown provide " + providerName)
	}
}

type memoryProvider struct {
	sync.Map
}


func GenerateSessionID() string {
	return uuid.Must(uuid.NewV4()).String()
}
func (m *memoryProvider) SessionInit(sid string)(*Session, error) {
	expiresAt := time.Now().Add(24 * time.Hour)
	session := &Session {
		ID: sid,
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

func (m *memoryProvider) SessionDestroy(sid string) error {
	m.Delete(sid)
	return nil
}

func (m *memoryProvider) SessionGC() {
	for{
		time.Sleep(time.Minute * 10)
		m.Range(func(key, value interface{}) bool {
			if time.Now().UnixNano() >= value.(*Session).ExpiresAt.UnixNano() {
				m.Delete(key)
			}
			return true
		})
		runtime.GC()
	}
}

func (m *memoryProvider) SessionUpdate(key string, value interface {})(error) {
	m.Store(key, value)
	return nil
}

var globalManager *Manager

func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session *Session){
	cookie, err := r.Cookie("session_id")
	var sess *Session
	if err != nil || cookie.Value == "" {
		sessionId := GenerateSessionID()
		fmt.Println(1)
		sess, _ = m.provider.SessionInit(sessionId)
		fmt.Println(2)
		fmt.Println(sess.ID)
		cookie := http.Cookie{ Name: "session_id", Value: session.ID, Path: "/", HttpOnly: true, MaxAge: int(m.maxLifeTime)}
		fmt.Println(3)
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		sess, _= m.provider.SessionRead(sid)
	}
	return  sess
}

func (s *Session) Get(key string) (*Session) {
	sess, err :=globalManager.provider.SessionRead(key)
	if err != nil {
		return sess
	} else {
		sess, _ :=	globalManager.provider.SessionInit(key)
		return sess
	}
}

func (s *Session) Set(key string, value interface{}) error {
	// globalManager.provider.SessionInit()
	globalManager.provider.SessionUpdate(key, value)
	return nil
}


// TODO session 还需要提供set get方法
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalManager.SessionStart(w, r)
	fmt.Print(sess)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("session3/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	
	cookie, err := r.Cookie("session_id")
	fmt.Println(cookie.Value)
	if cookie.Value == "" || err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		fmt.Fprintf(w, "Welcome, User!")
	}

}

func main() {
	NewProvider("memory")
	globalManager, _ = NewManager("memory", 10)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8081", nil)
}


