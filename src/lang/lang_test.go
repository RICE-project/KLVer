package lang

import "testing"


func TestReadLang(t *testing.T){
        zhLang, err1 := ReadLang("zh_CN")
        if err1 != nil {
                t.Errorf("ERR:",err1)
        }
        t.Logf("DEB:", zhLang)
        enLang, err2 := ReadLang("en")
        if err2 != nil {
                t.Errorf("ERR:",err2)
        }
        t.Logf("DEB:", enLang)
        if zhLang["TEST"] != "这是一条测试文本。ABC123" {
                t.Errorf("ERR: %s", zhLang["TEST"])
        }
        if enLang["TEST"] != "This is a test message.ABC123" {
                t.Errorf("ERR: %s", enLang["TEST"])
        }
}
