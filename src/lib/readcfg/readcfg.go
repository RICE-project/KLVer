/*
Common library of config and lang module.
*/
package readcfg

import "io/ioutil"
import "strings"

/*
ReacConfig is used to read text config file.

any config should be writen like this.

    key1 = value1
    key2 = value2
    key3 = value3
    ...

Blank line(s) and line(s) started with '#' will be skiped.
*/
func ReadConfig(filename string) (map[string]string, error) {
	cfg := make(map[string]string)
	b, err := ioutil.ReadFile(filename)
	var key, value string
/*	j := 0
	line := make([]byte, 5)
	for i, s := range b {
		if s == 10 {
			line = b[j:i]
			j = i + 1
		} else {
			continue
		}
		if len(line) == 0 || line[0] == '#' {
			continue
		} //Blank line and start with '#' will be skiped
		for k, t := range line {
			if t == '=' {
				key = strings.Replace(string(line[:k])
				value = string(line[k+1:])
				break //No need go on.
			}
		}
		cfg[key] = value
	}
	return cfg, err
    */
    lines := strings.Split(string(b), "\n")
    lineView := make([]string, 2)
    for _, line := range(lines) {
        line = strings.Replace(line, "\r", "", -1)
        line = strings.Replace(line, " ", "", -1)
        if len(line) == 0 || line[0] == '#' {
            continue
        }
        lineView = strings.Split(line, "=")
        key = lineView[0]
        value = lineView[1]
        cfg[key] = value
    }
    return cfg, err
}
