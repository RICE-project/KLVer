package sessions

import "testing"
import "time"

var data map[string]interface{} = map[string]interface{}{"name": "Admin", "role": 1, "age": "23", "isAdmin": "0"}

func TestTypeSession(t *testing.T) {
	a := new(Session)
	a.newSession(data)
	t.Log("Session:", a)

	t.Log("ExpireTime: ", a.getExpireTime())
	t.Log("Sid: ", a.GetSid())

	value := a.GetValue()

	for key, val := range value {
		t.Log("Key: ", key, "\tValue:", val)
	}
}

func TestTypeSessionManager(t *testing.T) {
	a := new(SessionManager)
	a.Init()
	b := a.CreateSession(data)
	t.Log("Session Manager: ", a)
	sid := b.GetSid()
	t.Log("SID: ", sid)
	t.Log("Session: ", b)
	data = map[string]interface{}{"name": "Test"}
	b.SetValue(data)
	time.Sleep(3 * time.Second)
	c, _ := a.GetSession(sid)
	t.Log("Session: ", c)
}
