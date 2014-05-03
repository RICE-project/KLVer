//HTML render here!
package page

import (
	"html/template"
	"io/ioutil"
	"lib/consts"
	"lib/logger"
	"lib/sessions"
    "lib/readcfg"
	"net/http"
	"path"
	"strings"
    "fmt"
)

const listDir = 0x0001

type Page struct {
	lang         *map[string]string
	sessionM     *sessions.SessionManager
	templates    map[string]*template.Template
	templatesErr map[string]*template.Template
	log          *logger.Logger
    mimeType     map[string]string
}

//Load language and HTML templates.
func (this *Page) Init(language *map[string]string, sessionManager *sessions.SessionManager, logs *logger.Logger) error {
    var errMime error
	this.SetLang(language)
    this.mimeType, errMime = getMimetype()
    if errMime != nil {
        this.log.LogCirtical(errMime)
        return errMime
    }
	this.sessionM = sessionManager
	this.log = logs
	this.templates = make(map[string]*template.Template)
	this.templatesErr = make(map[string]*template.Template)
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

func (this *Page) cachePage(dir string) error {
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
		templateNameShort := strings.TrimSuffix(templateName, ".html") //No extision name.
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
		} else if name == "" { // Default page.
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
	return func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL.Path
		urla := strings.Split(url, "/")
		file := staticDir + urla[len(urla)-1] // real file path.
		//HACK: for svg
		//writer.Header().Set("Content-Type", hackHeader(url)+"; charset=utf-8")
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

func (this *Page) SetLang(language *map[string]string) {
	this.lang = language
}

func (this *Page) setMimeType(writer *http.ResonseWriter, url string){
    ext := path.Ext(url)
    mimeType, found := this.mimeType[ext]
    if found {
        (*writer).Header().set("Content-Type", fmt.Sprintf("%s; charset=utf-8", mimeType))
    }
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//HACK: for incorrect mime-type (such as svg)
func hackHeader(url string) string {
	ext := path.Ext(url)
	mimeType := ""
	switch ext {
	case ".svg":
		mimeType = "image/svg+xml"
	case ".css":
		mimeType = "text/css"
	case ".js":
		mimeType = "text/javascript"
	}
	return mimeType
}

func getMimetype() (map[string]string,error){
    mime, err := readcfg.ReadConfig(consts.DIR_CFG + consts.CFG_MIMETYPE_FILE)
    return mime, err
}
