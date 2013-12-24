
package sessions

import(
        "sync"
        "time"
        "lib/consts"
	"errors"
)

//A global Session Manager.
type SessionManager struct{
        lock sync.Mutex
        sessionList map[string] *Session
}

//Init the Session Manager.
func (this *SessionManager) Init(){
        this.sessionList = make(map[string] *Session)
        go this.gc()
}

//Creat a new Session.
func (this *SessionManager) CreateSession(value map[string] string) *Session {
        b := new(Session)
        b.newSession(value)
        sid := b.GetSid()
        this.sessionList[sid] = b
        return b
}

//Get a Session.
func (this *SessionManager) GetSession(sid string) (*Session, error){
        ses, ok := this.sessionList[sid]
	if ok {
		ses.updateExpireTime()
		return ses, nil
	} else {
		return ses, errors.New("No this session")
	}
}

//Logout or time expired.
func (this *SessionManager) DestorySession(sid string) {
        delete(this.sessionList, sid)
}

//Sessions which are time-expired should be deleted.
func (this *SessionManager) gc() {
        gcList := make([]string, 0)
        this.lock.Lock()
        defer this.lock.Unlock()
        if len(this.sessionList) !=0 {
                for key, value := range this.sessionList {
                        if value.isExpired() {
                                gcList = append(gcList, key)
                        }
                }
                for _, gcSid := range gcList {
                        this.DestorySession(gcSid)
                }
        }
        time.AfterFunc(consts.CFG_GC_INTERVAL * time.Minute, func() {this.gc()})
}
