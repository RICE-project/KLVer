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

func (s *Session) newSession(value map[string]interface{}) {
	s.sid = generateSid()
	s.expire = expireTime()
	s.SetValue(value)
}

//Get Session values.
func (s *Session) GetValue() map[string]interface{} {
	return s.values
}

//Get Session ID.
func (s *Session) GetSid() string {
	return s.sid
}

func (s *Session) isExpired() bool {
	return s.expire < time.Now().Unix()
}

func (s *Session) getExpireTime() int64 {
	return s.expire
}

func (s *Session) updateExpireTime() {
	s.expire = expireTime()
}

//Set Session values.
func (s *Session) SetValue(value map[string]interface{}) {
	s.values = value
}
