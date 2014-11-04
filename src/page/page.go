//HTML render here!
package page

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"lib/consts"
	"lib/logger"
	"lib/readcfg"
	"lib/sessions"
	"net/http"
	"path"
	"strings"
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

	this.log = logs
	this.SetLang(language)
	this.mimeType, errMime = getMimetype()
	if errMime != nil {
		this.log.LogCritical(errMime)
		return errMime
	}

	this.sessionM = sessionManager
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
		defer func() {
			x := recover()
			if x != nil {
				file = staticDir
				this.err404Handler(writer, request)
			}
		}()

		this.setMimeType(&writer, url)
		this.log.LogInfo("HTTP GET:", url, "\t| staticDir:", staticDir, "\tfileName:", file)
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

func (this *Page) ForwardToHTTPS(httpPort string, httpsPort string, log *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		location := "https://" + strings.Replace(request.Host, ":"+httpPort, "", -1) + ":" + httpsPort + request.URL.Path
		log.LogInfo(location)
		writer.Header().Add("Location", location)
		writer.WriteHeader(http.StatusMovedPermanently)
	}
}

func (this *Page) setMimeType(writer *http.ResponseWriter, url string) {
	ext := path.Ext(url)
	mimeType, found := this.mimeType[ext[1:]] // skip the dot '.'
	if found {
		(*writer).Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", readcfg.TrimString(mimeType)))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getMimetype() (map[string]string, error) {
	mime, err := readcfg.ReadConfig(consts.DIR_CFG + consts.CFG_MIMETYPE_FILE)
	return mime, err
}
