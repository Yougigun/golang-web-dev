package session

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protect session
	provider    Provider
	maxLifeTime int64
}

func NewManager(providerName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[providerName]

	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}

	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

// Unique Session ID
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(
	w http.ResponseWriter, r *http.Request) (session Session) {
	// may cause performance problem
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session = manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:  manager.cookieName,
			Value: url.QueryEscape(sid),
			Path:  "/", HttpOnly: true,
			MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// Destroy Session ID
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{
			Name: manager.cookieName,
			Path: "/", HttpOnly: true, 
			Expires: expiration, 
			MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (manager *Manager) CookieExist(req *http.Request) (bool) {
	_, err := req.Cookie(manager.cookieName)
	if err!=nil{
		return false
	}
	return true
}

// Garbage Collection 
func (manager *Manager) GC() {
    manager.lock.Lock()
    defer manager.lock.Unlock()
    manager.provider.SessionGC(manager.maxLifeTime)
    time.AfterFunc(time.Duration(manager.maxLifeTime)*time.Second, func() { manager.GC() })
}

