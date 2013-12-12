package sessions

import(
        "lib/consts"
        "time"
        "crypto/md5"
        "encoding/hex"
        "strconv"
        "errors"
)

type session struct {
        name string
        exprreTime int64  //timestamp
}

type Sessions map[string] session

func (a *Sessions) Init() error {
        a = make(Sessions)
}

func (a *Sessions) SetSession(name string) {
        timeNow = time.Now().Unix()
        m := md5.New()
        m.Write([]byte(name + strconv.Itoa(timeNow)))
        sid := hex.EncodeToString(m.Sum(nil))
        a[sid] = session{name, (timeNow + consts.CFG_SESSION_TIMEOUT)}
}

func (a *Sessions) GetName(sid string) (string, error) {
        ses, ok := a[sid]

        if ok {
                err := nil
        } else {
                err := errors.New("No such session ID.")
        }

        return ses.name, err
}

func (a *Sessions) gc() {
        var ses session
        if len(a) != 0{
                for 
        }
}
