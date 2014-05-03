/*
Common library of config and lang module.
*/
package readcfg

import "io/ioutil"
import "strings"

/*
ReacConfig is used to read text config file.

any config should be writen like this.

    key1=value1
    key2=value2
    key3=value3
    ...

Blank line(s) and line(s) started with '#' will be skiped.

*/
func ReadConfig(filename string) (map[string]string, error) {
	var key, value, lineTest string
	cfg := make(map[string]string)
	b, err := ioutil.ReadFile(filename)
    cfgs := make([]string, 2)
    lines := strings.Split(
        strings.Replace(string(b), "\r", "", -1),  // Get rid of \r
        "\n")
    for _, line := range(lines) {
        lineTest = strings.Replace(line, " ", "", -1)
        lineTest = strings.Replace(lineTest, "\t", "", -1)
        if len(lineTest) == 0 || line[0] == '#' {
            continue
        } // Blank line and start with '#' will be skiped.
        cfgs = strings.SplitN(line, "=", 2)
        key = strings.Replace(cfgs[0], " ", "", -1)
        key = strings.Replace(key, "\t", "", -1)
        value = cfgs[1]
        cfg[key] = value
    }
	return cfg, err
}

func TrimString(str string)string{
    out := strings.Replace(str, " ", "", -1)
    out = strings.Replace(out, "\t", "", -1)
    return out
}
