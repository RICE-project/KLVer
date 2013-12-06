package lang

import "io/ioutil"
import "consts"

func ReadLang(language string) (map[string] string,error){
        lang := make(map[string] string)
        b, err := ioutil.ReadFile(consts.DIR_LANG + language + ".lang")
        if err != nil { panic(err) }
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
                print(key, value)
                lang[key] = value
        }
        return lang, nil  //No errors.
}
