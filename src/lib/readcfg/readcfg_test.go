package readcfg

import "testing"
import "lib/consts"

func TestReadConfig(t *testing.T) {
	test, err := ReadConfig("../" + consts.DIR_CFG + "klver.cfg")
	if err != nil {
		t.Error("Err: ", err)
	}
	if test["lang"] != "zh_CN" {
		t.Errorf("Err: ReadConfig() failed. expected 'zh_CN', got %s", test["lang"])
	}
}
