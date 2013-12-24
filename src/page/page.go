//HTML render here!
package page

import(
        "io/ioutil"
        "lib/consts"
        "net/http"
        "html/template"
        "path"
	"lib/sessions"
	"lib/logger"
)

type Page struct {
        lang *map[string] string
	sessionM *sessions.SessionManager
        templates map[string] *template.Template
	log *logger.Logger
}

//Load language and HTML templates.
func (this *Page) Init(language *map[string] string, sessionManager *sessions.SessionManager, logs *logger.Logger) error{
        this.SetLang(language)
	this.sessionM = sessionManager
	this.log = logs
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
		this.log.LogInfo("Loading HTML template '", templateName, "' done.")
        }
}

//Return http handler.
func (this *Page) GetHandler(name string, checkSession bool = consts.CHECK_SESSION_NO) http.HandlerFunc {
        return func(writer http.ResponseWriter, request *http.Request) {
		//TODO: if checkSession = true go login
		if checkSession {
			_,errSession := this.sesssionM.GetSession()
			if errSession != nil {
				errT := this.templates[consts.HTTP_LOGIN].Execute(writer, this.lang)
			}
		}
                err := this.templates[name].Execute(writer, this.lang)
                checkErr(err)
        }
}

//Return Templates List.
func (this *Page) GetTemplatesList() []string {
	templatesList := make([]string)
	for key, _ := range this.templates {
		templatesList = append(templatesList, key)
	}
	return templatesList
}

func (this *Page) SetLang(language *map[string] string) {
	this.lang = language
}

func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}
