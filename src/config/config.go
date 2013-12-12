package config

import (
        "lib/consts"
        "lib/readcfg"
)

type Config struct{
        Cfg map[string] string
}

func (a *Config) Init() error{
        var err error
        a.Cfg, err = readcfg.ReadConfig(consts.DIR_CFG + "glvsadm.cfg")
        return err  //No errors.
}

func (a *Config) GetDSN() string {
        return a.Cfg["mysql_user"] + ":" + a.Cfg["mysql_password"] + "@tcp(" + a.Cfg["mysql_host"] + ":" + a.Cfg["mysql_port"] + ")/" + a.Cfg["mysql_db"] + "?charset=utf8"
}
