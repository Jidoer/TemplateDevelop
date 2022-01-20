/**
 * @Author: FLYKO
 * @Description:
 * @File: JomePageRuner
 * @Version: 1.0.0
 * @LOVE AChu FOREVER
 * @Date: 2021/07/31 19:49
 * @本程序专门用于JomeTemplate开发
 * @Website: www.jomeidc.com
 */

package main

import (
	"2-JomeIDC/file"
	"encoding/json"
	"github.com/kataras/iris/v12"
)

type picture struct {
	ID      string
	URL     string
	Subject string
}

func main() {
	app := iris.New()
	tmpl := iris.Django("./Template", ".html")
	tmpl.Reload(false)
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings, " + s + "!"
	})
	app.RegisterView(tmpl)

	/************************************************************************
	 * ///////////////////////////////////////////////////////////////////////
	 *                JomePageRuner V 1.0 Powered By FlyKO
	 * ///////////////////////////////////////////////////////////////////////
	 ************************************************************************/

	app.Get("/{template:string}", func(ctx iris.Context) {
		t := ctx.Params().Get("template")
		app.Logger().Info(t) 
		pics := file.Reader("language/CN-zh.json")
		// JSONToMap
		var tempMap map[string]interface{}
		err := json.Unmarshal([]byte(pics), &tempMap)
		if err != nil {
			panic(err)
		}
		ctx.ViewData("language",tempMap)
	    ctx.View(t + ".html")
	})
	app.Get("/panel/{template:string}", func(ctx iris.Context) {

		//SYSTEM VAR






		t := ctx.Params().Get("template")
		app.Logger().Info(t) 
		pics := file.Reader("language/CN-zh.json")
		// JSONToMap
		var tempMap map[string]interface{}
		err := json.Unmarshal([]byte(pics), &tempMap)
		if err != nil {
			panic(err)
		}
		ctx.ViewData("language",tempMap)
	    ctx.View( "panel/"+ t +".html")
	})

	app.Favicon("./static/favicon.ico")
	//INDEX Page STATIC
	app.HandleDir("/js", "./static/js")
	app.HandleDir("/images","./static/images")
	app.HandleDir("/css","./static/css")
	app.HandleDir("/fonts","./static/fonts")
	app.HandleDir("/login2","./Server")

	//IDC PANEL STATIC
	app.HandleDir("/panel/js", "./Template/panel/panel/js")
	app.HandleDir("/panel/images", "./Template/panel/panel/images")
	app.HandleDir("/panel/fonts","./Template/panel/panel/fonts")
	app.HandleDir("/panel/css","./Template/panel/panel/css")

	app.Run(iris.Addr(":8899"))
	//app.Run(iris.TLS("0.0.0.0:443","m.crt","m.key"))

}


