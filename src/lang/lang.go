//Implement of l18n.
package lang

import (
	"lib/consts"
	"lib/readcfg"
)

//Get a language config.
func ReadLang(language string) (map[string]string, error) {
	lang, err := readcfg.ReadConfig(consts.DIR_LANG + language + ".lang")
	return lang, err //No errors.
}
