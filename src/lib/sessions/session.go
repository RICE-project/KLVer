/*
A implement of Session data structure.
*/
package sessions

import (
	"crypto/md5"
	"encoding/hex"
	"lib/consts"
	"math/rand"
	"strconv"
	"time"
)

type Session struct {
	sid    string
	expire int64                  //timestamp
	values map[string]interface{} //easy to expand
}

func generateSid() string {
	combo := make([]byte, 0)
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

func (this *Session) newSession(value map[string]interface{}) {
	this.sid = generateSid()
	this.expire = expireTime()
	this.SetValue(value)
}

//Get Session values.
func (this *Session) GetValue() map[string]interface{} {
	return this.values
}

//Get Session ID.
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

//Set Session values.
func (this *Session) SetValue(value map[string]interface{}) {
	this.values = value
}
