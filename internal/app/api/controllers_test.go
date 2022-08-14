package api

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

func Test_Health(t *testing.T) {
	port := ":1234"
	host := "http://localhost"
	serv := server{
		serv: &fasthttp.Server{},
		port: port,
	}
	serv.serv.Handler = serv.configureRouter().Handler

	defer startServerOnPort(t, port, serv.health).Close()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(host + port + "/health")
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")

	resp := fasthttp.AcquireResponse()

	assert.Nil(t, fasthttp.Do(req, resp))

	if resp.StatusCode() != 200 {
		t.Error("Error status code")
	}
	if len(resp.Body()) != 0 {
		t.Error("Error Body response")
	}
}

func TestNewServer(t *testing.T) {
	assert.NotNil(t, NewServer(""))
}

func TestServer_Run(t *testing.T) {
	serv := NewServer("")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		assert.Nil(t, serv.Run(ctx))
	}()
	cancel()
	time.Sleep(1 * time.Second)
}

func startServerOnPort(t *testing.T, port string, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost%s", port))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %s: %s", port, err)
	}

	go func() {
		err = fasthttp.Serve(ln, h)
		if err != nil {
			log.Println(err)
		}
	}()

	return ln
}
