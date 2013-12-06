package readcfg

import "io/ioutil"
import "strconv"

type Config struct{
        HttpPort int
        MysqlDb string
        MysqlUser string
        MysqlPass string
        MysqlHost string
        MysqlPort int
        Language string
}

func (a *Config) ReadConfig(filename string) error{
        b, err := ioutil.ReadFile(filename)
        if err != nil { return err }
        j := 0
        line := make([]byte, 5)
        var key, value string
        for i, v := range b {
                if v == 10 {
                        line = b[j:i]
                        j = i + 1
                } else {
                        continue
                }
                if len(line) == 0 || line[0] == 35 { continue }  //Start with '#' will be skiped
                for k, w := range line {
                        if w == 61 {
                                key = string(line[:k])
                                value = string(line[k + 1:])
                                break  //No need go on.
                        }
                }
                switch key{
                case "http_port":
                        a.HttpPort, _ = strconv.Atoi(value)
                case "mysql_host":
                        a.MysqlHost = value
                case "mysql_db":
                        a.MysqlDb = value
                case "mysql_user":
                        a.MysqlUser = value
                case "mysql_password":
                        a.MysqlPass = value
                case "mysql_port":
                        a.MysqlPort, _ = strconv.Atoi(value)
                case "lang":
                        a.Language = value
               }
        }
        return nil  //No errors.
}

func (a *Config) GetDSN() string {
        return a.MysqlUser + ":" + a.MysqlPass + "@tcp(" + a.MysqlHost + ":" + strconv.Itoa(a.MysqlPort) + ")/" + a.MysqlDb + "?charset=utf8"
}
