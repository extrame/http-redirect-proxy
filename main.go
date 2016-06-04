package main

import (
	"config"
	"flag"
	"fmt"
	"os"

	"github.com/bbangert/toml"
	"github.com/valyala/fasthttp"
	// "github.com/go-xorm/core"
	// _ "github.com/go-sql-driver/mysql"
	"net/http"
)

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	host := string(ctx.Host())
	url := string(ctx.RequestURI())
	redirectHost := config.Redirect[host]
	ctx.SetStatusCode(http.StatusTemporaryRedirect)
	ctx.Response.Header.Set("Location", fmt.Sprintf("http://%s%s", redirectHost, url))
}

func main() {
	//////////读取配置文件////////////////
	flag_conf := flag.String("config", "proxy.conf", "配置文件地址")
	flag.Parse()

	conf := new(config.Config)

	if _, err := toml.DecodeFile(*flag_conf, conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fasthttp.ListenAndServe(fmt.Sprintf(":%d", conf.Basic.Port), fastHTTPHandler)
}

// pass plain function to fasthttp
