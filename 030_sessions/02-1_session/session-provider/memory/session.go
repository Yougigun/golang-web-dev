package memory

import (
	"log"
	"sync"
)


type session struct {
	sync.RWMutex
	data map[string]interface{}
}

func (s *session) Set(key string, value interface{}) error {
	s.data[key] = value
	return nil
}

// Because session is map and map is reference type for parameter
func (s *session) Get(key string) (interface{}, error) {
	v := s.data[key]
	return v, nil
}

func (s *session) Delete(key string) error {
	delete(s.data, key)
	return nil
}

func (s *session) SessionID() string {
	v, ok := s.data["sessionID"]
	if !ok {
		log.Println("Session withour ID")
	}
	return v.(string)
}

func (s *session) Lock() {
	s.RWMutex.Lock()
}

func (s *session) Unlock() {
	s.RWMutex.Unlock()
}

func (s *session) RLock(){
	s.RWMutex.RLock()
}

func (s *session) RUnlock(){
	s.RWMutex.RUnlock()
}
