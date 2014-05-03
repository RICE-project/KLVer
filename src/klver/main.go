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
	httpPort, errPort := cfg.GetConfig("http_port")
	if errPort != nil {
		log.LogInfo("No http_port found in config file, use :80")
		httpPort = "80"
	}
	log.LogInfo("HTTP Serve at :", httpPort)
	errHttp := http.ListenAndServe(":" + httpPort, mux)
	if errHttp != nil {
		log.LogCritical(errHttp)
	}
}
