package sessions

import (
	"errors"
	"lib/consts"
	"lib/logger"
	"sync"
	"time"
)

//A global Session Manager.
type SessionManager struct {
	lock        sync.Mutex
	sessionList map[string]*Session
	log         *logger.Logger
}

//Init the Session Manager.
func (s *SessionManager) Init(log *logger.Logger) {
	s.sessionList = make(map[string]*Session)
	s.log = log
	go s.gc()
}

//Creat a new Session.
func (s *SessionManager) CreateSession(value map[string]interface{}) *Session {
	b := new(Session)
	b.newSession(value)
	sid := b.GetSid()
	s.sessionList[sid] = b
	return b
}

//Get a Session.
func (s *SessionManager) GetSession(sid string) (*Session, error) {
	ses, ok := s.sessionList[sid]
	if ok {
		ses.updateExpireTime()
		return ses, nil
	} else {
		return ses, errors.New("No s session")
	}
}

//Logout or time expired.
func (s *SessionManager) DestorySession(sid string) {
	delete(s.sessionList, sid)
}

//Sessions which are time-expired should be deleted.
func (s *SessionManager) gc() {
	s.log.LogInfo("Session gc Start!")
	gcList := make([]string, 0)
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.sessionList) != 0 {
		for key, value := range s.sessionList {
			if value.isExpired() {
				gcList = append(gcList, key)
			}
		}
		for _, gcSid := range gcList {
			s.DestorySession(gcSid)
            s.log.LogInfo("Session ID=", gcSid," is expired.")
		}
	}
	time.AfterFunc(consts.CFG_GC_INTERVAL*time.Minute, func() { s.gc() })
}
