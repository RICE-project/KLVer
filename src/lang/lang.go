//Implement of l18n.
package lang

import (
	"lib/consts"
	"lib/readcfg"
    "fmt"
)

//Get a language config.
func ReadLang(language string) (map[string]string, error) {
	lang, err := readcfg.ReadConfig(consts.DIR_LANG + language + ".lang")
    setConst(&lang)
	return lang, err //No errors.
}

//WRITE INFOMATION INTO LANGUAGE
func setConst(lang *map[string]string) {
    (*lang)["HTML_VERINFO"] = fmt.Sprintf("%s %s", consts.NAME, consts.VERSION)
    (*lang)["HTML_AUTHOR"] = consts.AUTHOR
    (*lang)["HTML_AUTHOR_MAIL"] = consts.AUTHOR_MAIL
}
