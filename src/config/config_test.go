package config

import "testing"

var a Config

func TestReadConfig(t *testing.T) {
        err := a.Init()
        if err != nil {
                t.Error(err)
        }
        t.Log(a)
        s, _ := a.GetConfig("lang")
        if s != "zh_CN" {
                t.Errorf("Err: Init() failed. got %s, expected 'zh_CN'", s)
        }
}

func TestGetDSN(t *testing.T) {
        r := a.GetDSN()
        t.Log(r)
        if r != "glvsadm:glvsadm@tcp(localhost:3306)/glvsadm?charset=utf8" {
                t.Errorf("Err: GetDSN() failed. get %s", r)
        }
}
