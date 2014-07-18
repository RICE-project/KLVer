/*
Common library of config and lang module.
*/
package readcfg

import "io/ioutil"
import "strings"
import "regexp"

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

        // Test blank lines.
        lineTest = TrimString(line)
        if len(lineTest) == 0 || line[0] == '#' {
            continue
        } // Blank line and start with '#' will be skiped.

        cfgs = strings.SplitN(line, "=", 2)
        key = TrimString(cfgs[0])
        value = cfgs[1]
        cfg[key] = TrimString(value)
    }
	return cfg, err
}

// Get rid of spaces and tab.
func TrimString(str string)string{
    const TRIM_HEAD_SPACE = "^[ ]+"
    const TRIM_TAIL_SPACE = "[ ]+$"
    regHead, _ := regexp.Compile(TRIM_HEAD_SPACE)
    regTail, _ := regexp.Compile(TRIM_TAIL_SPACE)

    out := strings.Replace(str, "\t", "", -1)  // Go to hell the tab space.
    out = regHead.ReplaceAllLiteralString(out, "")
    out = regTail.ReplaceAllLiteralString(out, "")

    return out
}
