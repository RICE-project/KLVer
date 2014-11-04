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
    chHttp chan int
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
	language, errLang := lang.ReadLang(langset)
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
	mux := http.NewServeMux()

	//Static resource should be writen with "/" end.
	mux.HandleFunc("/js/", pag.GetStaticHandler(consts.DIR_JS))
	mux.HandleFunc("/css/", pag.GetStaticHandler(consts.DIR_CSS))
	mux.HandleFunc("/images/", pag.GetStaticHandler(consts.DIR_IMAGES))

	mux.HandleFunc("/", pag.GetHandler())
	//TODO: ajax

    //Try https.
    httpsPort, errPorts := cfg.GetConfig("https_port")
    if errPorts != nil {
        log.LogInfo("No https_port found in config file, use :443")
        httpsPort = "443"
    }

    useHttps, errUseHttps := cfg.GetConfig("use_https")
    isServeHttps := false
    if errUseHttps != nil {
        log.LogInfo("No use_https found in config file, disable HTTPS")
    }else if useHttps=="yes"{
        log.LogInfo("Try to use HTTPS")
        isServeHttps = true

        certFile, errCert := cfg.GetConfig("certfile")
        if errCert != nil{
            log.LogInfo("Error when loading SSL Certificate. disable HTTPS")
            isServeHttps = false
        }

        certKeyFile, errKey := cfg.GetConfig("certkeyfile")
        if errKey != nil{
            log.LogInfo("Error when loading SSL Certificate Key. disable HTTPS")
            isServeHttps = false
        }

        if isServeHttps{
            log.LogInfo("HTTPS Server at :", httpsPort)
            go servHttps(chHttps, log, httpsPort, certFile, certKeyFile, mux)
        }

    }else{
        log.LogInfo("HTTPS disabled")
    }


	httpPort, errPort := cfg.GetConfig("http_port")
	if errPort != nil {
		log.LogInfo("No http_port found in config file, use :80")
		httpPort = "80"
	}

	log.LogInfo("HTTP Serve at :", httpPort)
    if !isServeHttps{
        go servHttp(chHttp, log, httpPort, mux)
    }else{
        httpForward := http.NewServeMux()
        httpForward.HandleFunc("/", forwardToHttps(httpsPort))
    }

    <-chHttps
    <-chHttp
    log.LogInfo("Exit")
}

func servHttp(ch chan int, log *logger.Logger, httpPort string, mux *http.ServeMux){
    errHttp := http.ListenAndServe(":" + httpPort, mux)
    if errHttp != nil{
        log.LogError(errHttp)
    }
    ch<-1
    return
}

func servHttps(ch chan int, log *logger.Logger, httpsPort string, cert string, certkey string, mux *http.ServeMux){
    errHttps := http.ListenAndServeTLS(":" + httpsPort, cert, certkey, mux)
    if errHttps != nil{
        log.LogError(errHttps)
    }
    ch<-1
    return
}

func forwardToHttps(httpsPort string)http.HandlerFunc{
    return func(writer http.ResponseWriter, request *http.Request){
        location := "https://" + request.Host + ":" + httpsPort + "/"
        writer.Header().Set("Location", location)
    }
}
