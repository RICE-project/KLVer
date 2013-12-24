package main

import(
	"ajax"
	"page"
	"config"
	"lang"
	"page"
	"lib/sessions"
	"lib/db"
	"lib/logger"
	"lib/consts"
	"net/http"
)
func init() {
	cfg := new(config.Config)
	log := new(logger.Logger)
	ses := new(sessions.SessionManager)
	pag := new(page.Page)
	log.LogInfo("Initializing...")
}

func main() {
	log.SetNewLogger()
	log.LogInfo(consts.NAME, consts.VERISION)
	log.LogInfo("Reading config file...")
	cfg.Init()
	langset := cfg.GetValue("lang")
	log.LogInfo("Loading language file...", langset)
	language := lang.ReadLang(langset)
	log.LogInfo("Starting Session Manager...")
	ses.Init()
	log.LogInfo("Starting Page Manager...")
	pag.Init(&language, ses, log)
	log.LogInfo("Starting HTTP Service...")
	templates := pag.GetTemplatesList()
	http.HandleFunc("/", pag.GetHandler("main"))
	for _, t := range templates {
		http.HandleFunc("/" + t, pag.GetHandler(t))
	}
	errHttp := http.ListenAndServe(":" + cfg.GetValue("http_port"))
	if errHttp != nil {
		log.LogCritical(err)
	}
}
