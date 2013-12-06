package readcfg

type config struct{
        http_port int
        mysql_db string
        mysql_user string
        mysql_pass string
        mysql_host string
        mysql_port int
        language string
}

func (a config) ReadConfig(filename string) (*config, error){

}
