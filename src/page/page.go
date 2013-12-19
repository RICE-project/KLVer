//HTML render here!
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
        templates map[string] *template.Template
}

//Load language and HTML templates.
func (this *Page) Init(language *map[string] string) error{
        this.lang = language
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
                this.templates[templateName] = t
        }
}

//Return http handler.
func (this *Page) GetHandler(name string) http.HandlerFunc {
        return func(writer http.ResponseWriter, request *http.Request) {
                err := templates[name].Execute(writer, this.lang)
                checkErr(err)
        }
}

func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}
