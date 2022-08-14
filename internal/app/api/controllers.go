package api

import (
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func (s *server) health(ctx *fasthttp.RequestCtx) {
	log.Println("handle health")
	ctx.SetStatusCode(http.StatusOK)
}
