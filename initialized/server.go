package initialized

import (
	"accessToken/global"
	"net/http"
)

func RunServer() {
	router := NewRouters()
	global.GAL_Router = router
	port := global.GAL_Config.System.Port
	sv := &http.Server{
		Addr:    port,
		Handler: router,
	}
	err := sv.ListenAndServe()
	if err != nil {
		panic("Server Start Err:" + err.Error())
	}
}
