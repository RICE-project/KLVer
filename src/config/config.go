package config

import (
        "lib/consts"
        "lib/readcfg"
)

type Config struct{
        Cfg map[string] string
}

func (this *Config) Init() error{
        var err error
        this.Cfg, err = readcfg.ReadConfig(consts.DIR_CFG + "glvsadm.cfg")
        return err  //No errors.
}

func (this *Config) GetDSN() string {
        return this.Cfg["mysql_user"] + ":" + this.Cfg["mysql_password"] + "@tcp(" + this.Cfg["mysql_host"] + ":" + this.Cfg["mysql_port"] + ")/" + this.Cfg["mysql_db"] + "?charset=utf8"
}
