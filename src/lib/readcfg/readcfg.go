package readcfg

import "io/ioutil"

func ReadConfig(filename string) (map[string] string,error){
        cfg := make(map[string] string)
        b, err := ioutil.ReadFile(filename)
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
                if len(line) == 0 || line[0] == 35 { continue }  //Blank line and start with '#' will be skiped
                for k, w := range line {
                        if w == 61 {
                                key = string(line[:k])
                                value = string(line[k + 1:])
                                break  //No need go on.
                        }
                }
                print(key, value)
                cfg[key] = value
        }
        return cfg, err
}
