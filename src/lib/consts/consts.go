/*
All global consts should be writen here
*/
package consts

//NOTE: Any dir path must end with "/"
const(
	dIR_PREFIX = "../../"
        DIR_CFG = dIR_PREFIX + "share/etc/"
        DIR_HTML = dIR_PREFIX + "share/html/"
        DIR_JS = dIR_PREFIX + "share/js/"
        DIR_CSS = dIR_PREFIX + "share/css/"
        DIR_MAKE = dIR_PREFIX + "share/make/"
        DIR_LANG = dIR_PREFIX + "share/lang/"
        DIR_LOG = dIR_PREFIX + "../share/log/"  //logger in src/lib
        CMD_DEPLOY = dIR_PREFIX + "share/libexec/deploy.sh"
        CFG_SESSION_TIMEOUT int64 = 3600  //1 hour
        CFG_GC_INTERVAL = 30  //30 minutes
	CHECK_SESSION_YES = true
	CHECK_SESSION_NO = false
	HTTP_LOGIN = "login"  //go login.
	VERISION = "0.0.0-初"
	NAME = "gLVSAdm"
)
