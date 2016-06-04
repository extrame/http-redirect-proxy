package controller

import "github.com/extrame/goblet"

type Index struct {
	goblet.Route  "/index"
	goblet.Render "html=index"
}

func (i *Index) Get(cx *goblet.Context) {
	cx.EnableCache()
	cx.RespondOK()
}
