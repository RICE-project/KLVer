package db

import "testing"
import "config"

var a config.Config

func TestConnect(t *testing.T) {
        err := a.Init()
        if err != nil {
                t.Errorf("Err: %s", err)
        }
        pdb, errdb := Connect(a.GetDSN())
        if errdb != nil {
                t.Errorf("Err: %s", errdb)
        }
        row := pdb.QueryRow("SELECT username FROM user_management")
        var name string
        errrow := row.Scan(&name)
        if errrow != nil {
                t.Errorf("Err: %s", errrow)
        }
        t.Logf("Deb: %s", name)
        if name != "admin" {
                t.Errorf("Err: unexpected value '%s'", name)
        }
}


