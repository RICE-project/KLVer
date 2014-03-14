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
        templatesErr map[string] *template.Template
	log *logger.Logger
}

//Load language and HTML templates.
func (this *Page) Init(language *map[string] string, sessionManager *sessions.SessionManager, logs *logger.Logger) error{
        this.SetLang(language)
	this.sessionM = sessionManager
	this.log = logs
	this.templates = make(map[string] *template.Template)
        this.templatesErr = make(map[string] *template.Template)
	errCPage := this.cachePage(consts.DIR_HTML)
	if errCPage != nil {
                this.log.LogCritical(errCPage)
		return errCPage
	}
        errEPage := this.cachePage(consts.DIR_HTML_ERROR)
        if errEPage != nil {
                this.log.LogCritical(errEPage)
                return errEPage
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

                templatePath = dir + templateName
                t := template.Must(template.ParseFiles(templatePath))
		templateNameShort := strings.TrimSuffix(templateName, ".html")  //No extision name.
                this.templates[templateNameShort] = t
		this.log.LogInfo("Load HTML template '", templateName, "' done.")
        }
	return nil
}

//Return http handler.
func (this *Page) GetHandler() http.HandlerFunc {
        return func(writer http.ResponseWriter, request *http.Request) {
                name := strings.TrimLeft(request.URL.Path, "/")
		template, found := this.templates[name]
                // Check if the page exists.
                if found {
                        err := template.Execute(writer, this.lang)
                        checkErr(err)
                } else if name == "" {  // Default page.
                        template = this.templates[consts.HTTP_DEFAULT]
                        err := template.Execute(writer, this.lang)
                        checkErr(err)
                } else {
                        this.err404Handler(writer, request)
                }
        }
}

//Return static contents.
func (this *Page) GetStaticHandler(staticDir string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request){
		url := request.URL.Path
		file := staticDir + strings.Split(url, "/")[2]
                this.log.LogInfo("HTTP GET:", url, "| staticDir:", staticDir, "fileName:", file)
		http.ServeFile(writer, request, file)
	}
}

func (this *Page) err404Handler(writer http.ResponseWriter, request *http.Request) {
        this.log.LogWarning("404 Not Found when access:", request.URL.Path)
        writer.Header().Set("Content-Type", "text/html; charset=utf-8")
        writer.WriteHeader(http.StatusNotFound)
        template := this.templates["404"]
        err := template.Execute(writer, this.lang)
        checkErr(err)
}

func (this *Page) SetLang(language *map[string] string) {
	this.lang = language
}

func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}


