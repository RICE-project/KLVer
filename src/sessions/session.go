package sessions

import (
        "lib/consts"
        "time"
        "crypto/md5"
        "math/rand"
        "strconv"
        "encoding/hex"
)

type Session struct {
        sid string
        expire int64  //timestamp
        values map[string] string  //easy to expand
}

func generateSid() string{
        combo := make([]byte,0)
        timeNow := time.Now().Unix()
        randNum := rand.Int63()

        combo = strconv.AppendInt(combo, timeNow, 10)
        combo = strconv.AppendInt(combo, randNum, 10)
        w := md5.New()
        w.Write(combo)

        result := w.Sum([]byte(""))
        return hex.EncodeToString(result)
}

func expireTime() int64 {
        timeNow := time.Now().Unix()
        return timeNow + consts.CFG_SESSION_TIMEOUT
}

func (this *Session) newSession(value map[string] string) {
        this.sid = generateSid()
        this.expire = expireTime()
        this.SetValue(value)
}

func (this *Session) GetValue() map[string] string{
        return this.values
}

func (this *Session) GetSid() string {
        return this.sid
}

func (this *Session) isExpired() bool {
        return this.expire < time.Now().Unix()
}

func (this *Session) getExpireTime() int64 {
        return this.expire
}

func (this *Session) updateExpireTime() {
        this.expire = expireTime()
}

func (this *Session) SetValue(value map[string] string) {
        this.values = value
}
