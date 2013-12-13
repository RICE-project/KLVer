package sessions

import(
        "sync"
        "time"
        "lib/consts"
)

type SessionManager struct{
        lock sync.Mutex
        sessionList map[string] *Session
}

func (this *SessionManager) Init(){
        this.sessionList = make(map[string] *Session)
        go this.gc()
}

func (this *SessionManager) CreateSession(value map[string] string) *Session {
        b := new(Session)
        b.newSession(value)
        sid := b.GetSid()
        this.sessionList[sid] = b
        return b
}

func (this *SessionManager) GetSession(sid string) *Session{
        ses := this.sessionList[sid]
        ses.updateExpireTime()
        return ses
}

func (this *SessionManager) DestorySession(sid string) {
        delete(this.sessionList, sid)
}

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
