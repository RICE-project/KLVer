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
	"strings"
)

const listDir = 0x0001

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
	this.templates = make(map[string] *template.Template)
	errCPage := this.cachePage(consts.DIR_HTML)
	if errCPage != nil {
		return errCPage
	}
	return nil
}

func (this *Page) cachePage(dir string) error{
        fileInfoArr, errReadDir := ioutil.ReadDir(dir)
        if errReadDir != nil {
                return errReadDir
        }

        var templateName, templatePath string
        for _, fileInfo := range fileInfoArr {
                templateName = fileInfo.Name()
                if ext := path.Ext(templateName); ext != ".html" {
			this.log.LogInfo("Skip non-template file: ", templateName)
                        continue
                }

                templatePath = consts.DIR_HTML + templateName
                t := template.Must(template.ParseFiles(templatePath))
		templateNameShort := strings.TrimSuffix(templateName, ".html")  //No extision name.
                this.templates[templateNameShort] = t
		this.log.LogInfo("Load HTML template '", templateName, "' done.")
        }
	return nil
}

//Return http handler.
func (this *Page) GetHandler(name string) http.HandlerFunc {
        return func(writer http.ResponseWriter, request *http.Request) {
		template := this.templates[name]
                err := template.Execute(writer, this.lang)
                checkErr(err)
        }
}

//Return static contents.
func (this *Page) GetStaticHandler(staticDir string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request){
		url := request.URL
		this.log.LogInfo("HTTP GET:", url, "staticDir", staticDir)
	}
}

//Return Templates List.
func (this *Page) GetTemplatesList() []string {
	templatesList := make([]string, 0)
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
