package model

import (
	"sync"
	"time"
)

//User struct
type User struct {
	Nombre   string
	Usuario  string
	Email    string
	Password string
}

//Foto struct
type Foto struct {
	Nombre string
	Texto  string
	URL    string
	Fecha  time.Time
}

//Comentario struct
type Comentario struct {
	Texto      string
	Comentario string
}

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0 significa que no esta especificado el atributo 'Max-Age'.
	// MaxAge<0 significa elimiar la cookie ahora mismo, equivalente a 'Max-Age: 0'
	// MaxAge>0 significa a cantidad de segundos que debe permanecer en memoria la cookie
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Texto plano con las parejas de atributos sin evaluar
}

//Manager struct
type Manager struct {
	cookieName  string     //private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

//Provider interface
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

//Session interface
type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}
