package session

var provides = make(map[string]Provider)

// Low Level Storage Structure. ex memory, file, etc. 
type Provider interface {
    SessionInit(sid string) (Session)
    SessionRead(sid string) (Session, error)
    SessionDestroy(sid string)
    SessionGC(maxLifeTime int64)
}

func Register(name string, provider Provider) {
    if provider == nil {
        panic("session: Register provider is nil")
    }
    if _, dup := provides[name]; dup {
        panic("session: Register called twice for provider " + name)
    }
    provides[name] = provider
}
