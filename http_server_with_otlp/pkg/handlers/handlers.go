package handlers

import (
	"errors"
	"net/http"

	"github.com/ralvescosta/gokit/httpw"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/internal/services"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type (
	HTTPHandlers interface {
		Install(router httpw.HTTPServer)
	}

	httpHandlers struct {
		service services.BookService
		logger  logging.Logger
		tracer  trace.Tracer
	}
)

func (h *httpHandlers) Install(router httpw.HTTPServer) {
	router.RegisterRoute(http.MethodPost, "/books", h.postHandler)
	router.RegisterRoute(http.MethodGet, "/books/{id}", h.getHandler)
	router.RegisterRoute(http.MethodGet, "/books", h.listHandler)
	router.RegisterRoute(http.MethodPut, "/books", h.putHandler)
	router.RegisterRoute(http.MethodDelete, "/books", h.deleteHandler)
}

func (h *httpHandlers) postHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("postHandler")
	h.service.RegisterBook(req.Context())
}

func (h *httpHandlers) getHandler(w http.ResponseWriter, req *http.Request) {
	_, span := h.tracer.Start(req.Context(), "getHandler")
	defer span.End()

	span.SetStatus(codes.Error, "some error")
	span.RecordError(errors.New("some error"))

	h.logger.Info("getHandler")
	h.service.GetBook(req.Context())
}

func (h *httpHandlers) listHandler(w http.ResponseWriter, req *http.Request) {
	_, span := h.tracer.Start(req.Context(), "listHandler")
	defer span.End()

	h.logger.Info("listHandler")
	h.service.ListBook(req.Context())

	w.Write([]byte("oi"))
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
	return &httpHandlers{service, logger, otel.Tracer("http handler")}
}
