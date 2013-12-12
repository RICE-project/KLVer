package lang

import (
        "lib/consts"
        "lib/readcfg"
)

func ReadLang(language string) (map[string] string,error){
        lang, err := readcfg.ReadConfig(consts.DIR_LANG + language + ".lang")
        return lang, err  //No errors.
}
