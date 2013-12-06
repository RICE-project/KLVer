package readcfg

import "testing"
import "consts"

func TestReadConfig(t *testing.T) {
        test, err := ReadConfig(consts.DIR_CFG + "glvsadm.cfg")
        if err != nil {
                t.Error("Err: ", err)
        }
        if test["lang"] != "zh_CN" {
                t.Errorf("Err: ReadConfig() failed. expected 'zh_CN', got %s", test["lang"])
        }
}
