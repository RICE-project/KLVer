/*
All global consts should be writen here
*/
package consts

//NOTE: Any dir path must end with "/"
const (
    dIR_PREFIX                = DIR + "share/"
    DIR_CFG                   = dIR_PREFIX + "etc/"
    DIR_HTML                  = dIR_PREFIX + "html/normal/"
    DIR_HTML_ERROR            = dIR_PREFIX + "html/error/"
    DIR_JS                    = dIR_PREFIX + "js/"
    DIR_CSS                   = dIR_PREFIX + "css/"
    DIR_IMAGES                = dIR_PREFIX + "images/"
    DIR_MAKE                  = dIR_PREFIX + "make/"
    DIR_LANG                  = dIR_PREFIX + "lang/"
    DIR_LOG                   = dIR_PREFIX + "log/" //logger in src/lib
    CMD_DEPLOY                = dIR_PREFIX + "libexec/deploy.sh"
    CFG_FILE                  = "klver.cfg"
    CFG_MIMETYPE_FILE         = "mimetype.cfg"
    CFG_SESSION_TIMEOUT int64 = 3600 //1 hour
    CFG_GC_INTERVAL           = 30   //30 minutes
    CHECK_SESSION_YES         = true
    CHECK_SESSION_NO          = false
    HTTP_LOGIN                = "login" //go login.
    HTTP_DEFAULT              = "main"  //index page.
    VERSION                   = "0.0.0-初"
    NAME                      = "KLVer"
    COOKIE_SESSION_NAME       = "GSESSION"
    AUTHOR                    = "TonyChyi"
    AUTHOR_MAIL               = "tonychee1989@gmail.com"
)
