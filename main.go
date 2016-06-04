package main

import (
	"fmt"
	"os"

	"github.com/extrame/http-redirect-proxy/config"

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
	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}
	fasthttp.ListenAndServe(fmt.Sprintf(":%s", port), fastHTTPHandler)
}

// pass plain function to fasthttp
