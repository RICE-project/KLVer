/*
Common library of config and lang module.
*/
package readcfg

import "io/ioutil"

/*
ReacConfig is used to read text config file.

any config should be writen like this.

    key1=value1
    key2=value2
    key3=value3
    ...

Blank line(s) and line(s) started with '#' will be skiped.

the '=' between key and value muse be without any space.
*/
func ReadConfig(filename string) (map[string] string,error){
        cfg := make(map[string] string)
        b, err := ioutil.ReadFile(filename)
        j := 0
        line := make([]byte, 5)
        var key, value string
        for i, s := range b {
                if s == 10 {
                        line = b[j:i]
                        j = i + 1
                } else {
                        continue
                }
                if len(line) == 0 || line[0] == 35 { continue }  //Blank line and start with '#' will be skiped
                for k, t := range line {
                        if t == 61 {
                                key = string(line[:k])
                                value = string(line[k + 1:])
                                break  //No need go on.
                        }
                }
                cfg[key] = value
        }
        return cfg, err
}
