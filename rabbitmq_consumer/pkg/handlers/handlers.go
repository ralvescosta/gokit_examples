package handlers

import (
	"net/http"

	"github.com/ralvescosta/gokit/httpw/server"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/internal/services"
)

type (
	HTTPHandlers interface {
		Install(router server.HTTPServer)
	}

	httpHandlers struct {
		service services.BookService
		logger  logging.Logger
	}
)

func (h *httpHandlers) Install(router server.HTTPServer) {
	router.Route(server.NewRouteBuilder().POST("/books").Handler(h.postHandler).Build())
	router.Route(server.NewRouteBuilder().GET("/books/${ID}").Handler(h.getHandler).Build())
	router.Route(server.NewRouteBuilder().GET("/books").Handler(h.listHandler).Build())
	router.Route(server.NewRouteBuilder().PUT("/books").Handler(h.putHandler).Build())
	router.Route(server.NewRouteBuilder().DELETE("/books").Handler(h.deleteHandler).Build())
}

func (h *httpHandlers) postHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("postHandler")
	h.service.RegisterBook(req.Context())
}

func (h *httpHandlers) getHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("getHandler")
	h.service.GetBook(req.Context())
}

func (h *httpHandlers) listHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("listHandler")
	h.service.ListBook(req.Context())
}

func (h *httpHandlers) putHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("putHandler")
	h.service.UpdateBook(req.Context())
}

func (h *httpHandlers) deleteHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("deleteHandler")
	h.service.DeleteBook(req.Context())
}

func NewHandler(logger logging.Logger, service services.BookService) HTTPHandlers {
	return &httpHandlers{service, logger}
}
