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
func (p *Page) Init(language *map[string]string, sessionManager *sessions.SessionManager, logs *logger.Logger) error {
	var errMime error

	p.log = logs
	p.SetLang(language)
	p.mimeType, errMime = getMimetype()
	if errMime != nil {
		p.log.LogCritical(errMime)
		return errMime
	}

	p.sessionM = sessionManager
	p.templates = make(map[string]*template.Template)
	p.templatesErr = make(map[string]*template.Template)
	errCPage := p.cachePage(consts.DIR_HTML)
	if errCPage != nil {
		p.log.LogCritical(errCPage)
		return errCPage
	}
	errEPage := p.cachePage(consts.DIR_HTML_ERROR)
	if errEPage != nil {
		p.log.LogCritical(errEPage)
		return errEPage
	}
	return nil
}

func (p *Page) cachePage(dir string) error {
	fileInfoArr, errReadDir := ioutil.ReadDir(dir)
	if errReadDir != nil {
		return errReadDir
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			p.log.LogInfo("Skip non-template file: ", templateName)
			continue
		}

		templatePath = dir + templateName
		t := template.Must(template.ParseFiles(templatePath))
		templateNameShort := strings.TrimSuffix(templateName, ".html") //No extision name.
		p.templates[templateNameShort] = t
		p.log.LogInfo("Load HTML template '", templateName, "' done.")
	}
	return nil
}

//Return http handler.
func (p *Page) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := strings.TrimLeft(request.URL.Path, "/")
		template, found := p.templates[name]
		// Check if the page exists.
		if found {
			err := template.Execute(writer, p.lang)
			checkErr(err)
		} else if name == "" { // Default page.
			template = p.templates[consts.HTTP_DEFAULT]
			err := template.Execute(writer, p.lang)
			checkErr(err)
		} else {
			p.err404Handler(writer, request)
		}
	}
}

//Return static contents.
func (p *Page) GetStaticHandler(staticDir string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL.Path
		urla := strings.Split(url, "/")
		file := staticDir + urla[len(urla)-1] // real file path.
		defer func() {
			a := recover()
			if a != nil {
				file = staticDir
				p.err404Handler(writer, request)
			}
		}()

		p.setMimeType(&writer, url)
		p.log.LogInfo("HTTP GET:", url, "\t| staticDir:", staticDir, "\tfileName:", file)
		http.ServeFile(writer, request, file)
	}
}

func (p *Page) err404Handler(writer http.ResponseWriter, request *http.Request) {
	p.log.LogWarning("404 Not Found when access:", request.URL.Path)
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusNotFound)
	template := p.templates["404"]
	err := template.Execute(writer, p.lang)
	checkErr(err)
}

func (p *Page) SetLang(language *map[string]string) {
	p.lang = language
}

func (p *Page) ForwardToHTTPS(httpPort string, httpsPort string, log *logger.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		location := "https://" + strings.Replace(request.Host, ":"+httpPort, "", -1) + ":" + httpsPort + request.URL.Path
		writer.Header().Add("Location", location)
		writer.WriteHeader(http.StatusMovedPermanently)
	}
}

func (p *Page) setMimeType(writer *http.ResponseWriter, url string) {
	ext := path.Ext(url)
	mimeType, found := p.mimeType[ext[1:]] // skip the dot '.'
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
