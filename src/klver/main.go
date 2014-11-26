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
	"os"
	"os/signal"
)

func main() {
	var err error

	cfg := new(config.Config)
	log := new(logger.Logger)
	ses := new(sessions.SessionManager)
	pag := new(page.Page)

	err = log.SetNewLogger()
	if err != nil {
		panic(err)
	}
	log.LogInfo("Initializing...")

	log.LogInfo(consts.NAME, consts.VERSION)
	log.LogInfo("Reading config file...")
	err = cfg.Init()
	if err != nil {
		log.LogCritical(err)
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
	err = pag.Init(&language, ses, log)
	if err != nil {
		log.LogCritical(err)
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

	chHttp := make(chan int)
	chHttps := make(chan int)
	httpPort, err := cfg.GetConfig("http_port")
	if err != nil {
		log.LogWarning("No http_port found in config file, use :80")
		httpPort = "80"
	}

	//Try https.
	log.LogInfo("Try to use HTTPS")
	useHttps, err := cfg.GetConfig("use_https")
	if err != nil {
		log.LogWarning("No use_https found in config file, disable HTTPS")
	}
	isServeHttps := (useHttps == "yes")

	//Load some staff.

	httpsPort, err := cfg.GetConfig("https_port")
	if err != nil {
		log.LogWarning("No https_port found in config file, use :443")
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
	chSignal := make(chan os.Signal, 1)
	go signal.Notify(chSignal)

	select {
	case <-chHttp:
		log.LogWarning("Thread HTTP exit.")
		os.Exit(1)
	case <-chHttps:
		log.LogWarning("Thread HTTPS exit.")
		os.Exit(1)
	case s := <-chSignal:
		log.LogWarning("Exit for signal", s)
	}
}

func servHttp(ch chan int, ser *http.Server ,log *logger.Logger, httpPort string, mux *http.ServeMux) {
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
