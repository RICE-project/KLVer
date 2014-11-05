/*
- KLVer By: Tony Chyi <tonychee1989@gmail.com> -

KLVer (pronounces 'clever') is a web-based keepalived configuration tool
developed with go programming language. It allows a user create/delete
LVS instance, add/remove MASTER/BACKUP server, add/remove virtual server
and realserver, and build configuration file to deploy them automatically.

This program is free software; you can redistribute it and/or modify it
under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 3, or (at your option)
any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the
Free Software Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301 USA.
*/

package main

import (
	//	"ajax"
	"config"
	"lang"
	"lib/sessions"
	"page"
	//	"lib/db"
	"lib/consts"
	"lib/logger"
	"net/http"
)

var (
	chHttps chan int
	chHttp  chan int
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

	log.LogInfo(consts.NAME, consts.VERSION)
	log.LogInfo("Reading config file...")
	errCfg := cfg.Init()
	if errCfg != nil {
		log.LogCritical(errCfg)
	}

	langset, _ := cfg.GetConfig("lang")
	log.LogInfo("Loading language file...", langset)
	language, err := lang.ReadLang(langset)
	if err != nil {
		log.LogCritical(err)
	}

	log.LogInfo("Starting Session Manager...")
	ses.Init(log)
	log.LogInfo("Starting Page Manager...")
	errPage := pag.Init(&language, ses, log)
	if errPage != nil {
		log.LogCritical(errPage)
	}
	log.LogInfo("Starting HTTP Service...")
	mux := http.NewServeMux()

	//Static resource should be writen with "/" end.
	mux.HandleFunc("/js/", pag.GetStaticHandler(consts.DIR_JS))
	mux.HandleFunc("/css/", pag.GetStaticHandler(consts.DIR_CSS))
	mux.HandleFunc("/images/", pag.GetStaticHandler(consts.DIR_IMAGES))
	mux.HandleFunc("/robots.txt", pag.GetStaticHandler(consts.DIR_MISC))
	mux.HandleFunc("/favicon.ico", pag.GetStaticHandler(consts.DIR_MISC))

	mux.HandleFunc("/", pag.GetHandler())
	//TODO: ajax

	httpPort, err := cfg.GetConfig("http_port")
	if err != nil {
		log.LogInfo("No http_port found in config file, use :80")
		httpPort = "80"
	}

	//Try https.
	log.LogInfo("Try to use HTTPS")
	useHttps, err := cfg.GetConfig("use_https")
	if err != nil {
		log.LogInfo("No use_https found in config file, disable HTTPS")
	}
	isServeHttps := (useHttps == "yes")

	//Load some staff.

	httpsPort, err := cfg.GetConfig("https_port")
	if err != nil {
		log.LogInfo("No https_port found in config file, use :443")
		httpsPort = "443"
	}

	certFile, err := cfg.GetConfig("certfile")
	if err != nil {
		log.LogWarning("No SSL certificate set, disable HTTPS")
		isServeHttps = false
	}

	certKeyFile, err := cfg.GetConfig("certkeyfile")
	if err != nil {
		log.LogWarning("No SSL certificate key set, disable HTTPS")
		isServeHttps = false
	}

	if isServeHttps {
		httpForward := http.NewServeMux()
		httpForward.HandleFunc("/", pag.ForwardToHTTPS(httpPort, httpsPort, log))
		go servHttp(chHttp, log, httpPort, httpForward)
		go servHttps(chHttps, log, httpsPort, certFile, certKeyFile, mux)
	} else {
		log.LogInfo("HTTPS disabled")
		go servHttp(chHttp, log, httpPort, mux)
	}

	<-chHttp
	<-chHttps

	log.LogInfo("Exit")
}

func servHttp(ch chan int, log *logger.Logger, httpPort string, mux *http.ServeMux) {
	log.LogInfo("HTTP Serve at :", httpPort)
	err := http.ListenAndServe(":"+httpPort, mux)
	if err != nil {
		log.LogWarning(err)
		ch <- 1
	}
}

func servHttps(ch chan int, log *logger.Logger, httpsPort string, cert string, certkey string, mux *http.ServeMux) {
	log.LogInfo("HTTPS Server at :", httpsPort)
	err := http.ListenAndServeTLS(":"+httpsPort, cert, certkey, mux)
	if err != nil {
		log.LogWarning(err)
		ch <- 1
	}
}
