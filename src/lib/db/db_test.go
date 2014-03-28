package db

import "testing"
import "config"

var a config.Config

func TestConnect(t *testing.T) {
	errCfg := a.Init()
	if errcfg != nil {
		t.Errorf("Err: %s", err)
	}
	pDb, errDb := Connect(a.GetDSN())
	if errDb != nil {
		t.Errorf("Err: %s", errDb)
	}
	row := pDb.QueryRow("SELECT username FROM user_management")
	var name string
	errRow := row.Scan(&name)
	if errRow != nil {
		t.Errorf("Err: %s", errRow)
	}
	t.Logf("Deb: %s", name)
	if name != "admin" {
		t.Errorf("Err: unexpected value '%s'", name)
	}
}
