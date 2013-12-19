//Read config file.
//
//config file is
//
//    <glvsadm_dir>/share/etc/glvsadm.cfg
//
package config

import (
        "lib/consts"
        "lib/readcfg"
        "fmt"
        "errors"
)

type Config struct{
        cfg map[string] string
}

//Init Config module.
func (this *Config) Init() error{
        var err error
        this.cfg, err = readcfg.ReadConfig(consts.DIR_CFG + "glvsadm.cfg")
        return err  //No errors.
}

//For database use.
func (this *Config) GetDSN() string {
        dsnTemplate := "%s:%s@tcp(%s:%s)/%s?charset=utf8"
        dsn := fmt.Sprintf(dsnTemplate, this.cfg["mysql_user"], this.cfg["mysql_password"], this.cfg["mysql_host"], this.cfg["mysql_port"], this.cfg["mysql_db"])
        return dsn
}

//Get a config value.
func (this *Config) GetConfig(key string) (string,error) {
        var err error
        config,ok := this.cfg[key]
        if !ok {
                err = errors.New("No such key in config file.")
        }
        return config, err
}
