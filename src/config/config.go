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
func (c *Config) Init() error {
	var err error
	c.cfg, err = ReadConfig(consts.DIR_CFG + consts.CFG_FILE)
	return err //No errors.
}

//For database use.
func (c *Config) GetDSN() string {
	const dsnTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8"
    mysqlUser, _ := c.GetConfig("mysql_user")
    mysqlPass, _ := c.GetConfig("mysql_password")
    mysqlHost, _ := c.GetConfig("mysql_host")
    mysqlPort, _ := c.GetConfig("mysql_port")
    mysqlDb, _ := c.GetConfig("mysql_db")
	return fmt.Sprintf(dsnTemplate, mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)
}

//Get a config value.
func (c *Config) GetConfig(key string) (string, error) {
	var err error
	config, ok := c.cfg[key]
	if !ok {
		err = errors.New("No such key in config file.")
	}
	return TrimString(config), err
}
