package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
)


type Session struct {
	ID        string
	Values    map[string]interface{}
	ExpiresAt time.Time
}

// SessionManager manages user sessions.
type SessionManager struct {
	sessions map[string]*Session
}

// NewSessionManager creates a new session manager.
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

// GenerateSessionID generates a new session ID.
func GenerateSessionID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// CreateSession creates a new session for the user.
func (sm *SessionManager) CreateSession() *Session {
	sessionID := GenerateSessionID()
	expiresAt := time.Now().Add(24 * time.Hour) // Session expires after 24 hours

	session := &Session{
		ID:        sessionID,
		Values:    make(map[string]interface{}),
		ExpiresAt: expiresAt,
	}

	sm.sessions[sessionID] = session
	return session
}

// GetSession retrieves a session based on the session ID.
func (sm *SessionManager) GetSession(sessionID string) *Session {
	session, ok := sm.sessions[sessionID]
	if !ok {
		return nil
	}

	// Check if the session has expired
	if session.ExpiresAt.Before(time.Now()) {
		delete(sm.sessions, sessionID)
		return nil
	}

	return session
}

// Middleware is a middleware handler that manages user sessions.
func (sm *SessionManager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			session := sm.CreateSession()
			cookie = &http.Cookie{
				Name:     "session_id",
				Value:    session.ID,
				Expires:  session.ExpiresAt,
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		} else if err == nil {
			sessionID := cookie.Value
			session := sm.GetSession(sessionID)
			if session == nil {
				// Session has expired or does not exist
				session = sm.CreateSession()
				cookie.Value = session.ID
				cookie.Expires = session.ExpiresAt
				http.SetCookie(w, cookie)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// Example usage
func main() {
	sessionManager := NewSessionManager()

	// Protected route that requires a valid session
	// Only accessible to users with a valid session
	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sessionID := cookie.Value
		session := sessionManager.GetSession(sessionID)
		if session == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// User is authenticated, handle protected content
		fmt.Fprintf(w, "Welcome, User!")
	})

	// Login route that creates a new session for the user
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		session := sessionManager.CreateSession()
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    session.ID,
			Expires:  session.ExpiresAt,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		fmt.Fprintf(w, "Logged in successfully!")
	})

	// Logout route that destroys the user's session
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sessionID := cookie.Value
		delete(sessionManager.sessions, sessionID)
		cookie.Expires = time.Now().Add(-time.Hour) // Expire the cookie
		http.SetCookie(w, cookie)

		fmt.Fprintf(w, "Logged out successfully!")
	})

	http.ListenAndServe(":8080", sessionManager.Middleware(http.DefaultServeMux))
}