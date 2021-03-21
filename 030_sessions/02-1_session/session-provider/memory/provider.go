package memory

import (
	"fmt"
	s "golang-web-dev/030_sessions/02-1_session/session"
	"time"
)

type provider struct {
	SessionDB map[string]*session
}

var p provider = provider{
	SessionDB: map[string]*session{},
}

func init() {
	s.Register("memory", &p)
}

func (p *provider) SessionInit(sid string) s.Session {
	// p.SessionDB[sid] = new(session)
	p.SessionDB[sid] = &session{data: make(map[string]interface{})}
	session := p.SessionDB[sid]
	session.Lock()
	defer session.Unlock()
	session.Set("sessionID", sid)
	session.Set("createdTime", time.Now())
	return session
}

func (p *provider) SessionRead(sid string) (s.Session, error) {
	session, ok := p.SessionDB[sid]
	if !ok {
		err := fmt.Errorf("access invalid session id %v", sid)
		return nil, err
	}
	return session, nil
}

func (p *provider) SessionDestroy(sid string)  {
	delete(p.SessionDB,sid)
}

func (p *provider) SessionGC(maxLifeTime int64) {
	now := time.Now()
	for sid := range p.SessionDB {
		session := p.SessionDB[sid]
		createdTime:=session.data["createdTime"].(time.Time)
		if now.After(createdTime.Add(time.Duration(maxLifeTime)*time.Second)) {
			delete(p.SessionDB,sid) 
		}
	}
}
