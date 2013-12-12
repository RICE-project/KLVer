package page

import(
        "io/ioutil"
        "lib/consts"
        "net/http"
        "html/template"
        "path"
)

type Page struct {
        lang *map[string] string
        tmpl map[string] *template.Template
}

//Load language and HTML templates.
func (a *Page) Init(language *map[string] string) error{
        a.lang = language
        fileInfoArr, err:= ioutil.ReadDir(consts.DIR_HTML)
        if errReadDir != nil {
                return err
        }

        var templateName, templelatePath string
        for _, fileInfo := range fileInfoArr {
                templateName = fileInfo.Name
                if ext := path.Ext(templateName); ext != "html" {
                        continue
                }

                templatePath = consts.DIR_HTML + templateName
                t := template.Must(template.ParseFiles(templatePath))
                a.tmpl[templateName] = t
        }
}

//Return http handler.
func (a *Page) GetHandler(name string) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                err := tmpl[name].Execute(w, a.lang)
                checkErr(err)
        }
}

func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}
