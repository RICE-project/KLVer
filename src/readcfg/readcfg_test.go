package readcfg

import "testing"
import "consts"

var a Config

func TestReadConfig(t *testing.T) {
        err := a.ReadConfig(consts.DIR_CFG + "glvsadm.cfg")
        if err != nil {
                t.Error(err)
        }
        t.Log(a)
        if a.Language != "zh_CN" {
                t.Errorf("Err: ReadConfig() failed. got %s, expected 'zh_CN'", a.Language)
        }
}

func TestGetDSN(t *testing.T) {
        r := a.GetDSN()
        t.Log(r)
        if r != "glvsadm:glvsadm@tcp(localhost:3306)/glvsadm?charset=utf8" {
                t.Errorf("Err: GetDSN() failed. get %s", r)
        }
}
