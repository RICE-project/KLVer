//Read config file.
//
//config file is
//
//    <klver_dir>/share/etc/klver.cfg
//
package config

import (
	"errors"
	"fmt"
	"lib/consts"
	. "lib/readcfg"
)

type Config struct {
	cfg map[string]string
}

//Init Config module.
func (this *Config) Init() error {
	var err error
	this.cfg, err = ReadConfig(consts.DIR_CFG + consts.CFG_FILE)
	return err //No errors.
}

//For database use.
func (this *Config) GetDSN() string {
	const dsnTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8"
    mysqlUser, _ := this.GetConfig("mysql_user")
    mysqlPass, _ := this.GetConfig("mysql_password")
    mysqlHost, _ := this.GetConfig("mysql_host")
    mysqlPort, _ := this.GetConfig("mysql_port")
    mysqlDb, _ := this.GetConfig("mysql_db")
	return fmt.Sprintf(dsnTemplate, mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)
}

//Get a config value.
func (this *Config) GetConfig(key string) (string, error) {
	var err error
	config, ok := this.cfg[key]
	if !ok {
		err = errors.New("No such key in config file.")
	}
	return TrimString(config), err
}
