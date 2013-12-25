package main

import(
//	"ajax"
	"page"
	"config"
	"lang"
	"lib/sessions"
//	"lib/db"
	"lib/logger"
	"lib/consts"
	"net/http"
)

func main() {
	cfg := new(config.Config)
	log := new(logger.Logger)
	ses := new(sessions.SessionManager)
	pag := new(page.Page)
	errLog := log.SetNewLogger()
	if errLog != nil {
		panic(errLog)
	}
	log.LogInfo("Initializing...")

	log.LogInfo(consts.NAME, consts.VERISION)
	log.LogInfo("Reading config file...")
	errCfg := cfg.Init()
	if errCfg != nil {
		log.LogCritical(errCfg)
	}

	langset, _ := cfg.GetConfig("lang")
	log.LogInfo("Loading language file...", langset)
	language, errLang:= lang.ReadLang(langset)
	if errLang != nil {
		log.LogCritical(errLang)
	}

	log.LogInfo("Starting Session Manager...")
	ses.Init(log)
	log.LogInfo("Starting Page Manager...")
	errPage := pag.Init(&language, ses, log)
	if errPage != nil {
		log.LogCritical(errPage)
	}
	log.LogInfo("Starting HTTP Service...")
	templates := pag.GetTemplatesList()
	mux := http.NewServeMux()

	mux.HandleFunc("/js", pag.GetStaticHandler(consts.DIR_JS))
	mux.HandleFunc("/css", pag.GetStaticHandler(consts.DIR_CSS))
	mux.HandleFunc("/image", pag.GetStaticHandler(consts.DIR_IMAGES))

	mux.HandleFunc("/", pag.GetHandler("main"))
	for _, t := range templates {
		mux.HandleFunc("/" + t, pag.GetHandler(t))
	}
	//TODO: ajax
	httpPort, errPort := cfg.GetConfig("http_port")
	if errPort != nil {
		httpPort = "80"
	}
	log.LogInfo("HTTP Serve at :", httpPort)
	errHttp := http.ListenAndServe(":" + httpPort, mux)
	if errHttp != nil {
		log.LogCritical(errHttp)
	}
}
