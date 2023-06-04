package httpHandlers

import (
	"errors"
	"net/http"

	"github.com/ralvescosta/gokit/httpw"
	"github.com/ralvescosta/gokit/logging"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	_ "github.com/ralvescosta/gokit_example/http_server_with_otlp/docs"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/internal/services"
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

	router.RegisterRoute(http.MethodGet, "/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3333/swagger/doc.json"),
	))
}

// CreateBook
// @Summary      CreateBook
// @Description  Create a new Book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param				 books	body		httpHandlers.Book	true	"Add Book"
// @Success      201  {object}  httpHandlers.Book
// @Failure      400  {object}  httpHandlers.HTTPError
// @Failure      404  {object}  httpHandlers.HTTPError
// @Failure      500  {object}  httpHandlers.HTTPError
// @Router       /books [post]
func (h *httpHandlers) postHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("postHandler")
	h.service.RegisterBook(req.Context())

	w.WriteHeader(200)
}

// GetBook
// @Summary      GetBook
// @Description  Get Book By ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  httpHandlers.Book
// @Failure      400  {object}  httpHandlers.HTTPError
// @Failure      404  {object}  httpHandlers.HTTPError
// @Failure      500  {object}  httpHandlers.HTTPError
// @Router       /books/{id} [get]
func (h *httpHandlers) getHandler(w http.ResponseWriter, req *http.Request) {
	_, span := h.tracer.Start(req.Context(), "getHandler")
	defer span.End()

	span.SetStatus(codes.Error, "some error")
	span.RecordError(errors.New("some error"))

	h.logger.Info("getHandler")
	h.service.GetBook(req.Context())

	w.WriteHeader(200)
}

func (h *httpHandlers) listHandler(w http.ResponseWriter, req *http.Request) {
	_, span := h.tracer.Start(req.Context(), "listHandler")
	defer span.End()

	h.logger.Info("listHandler")
	h.service.ListBook(req.Context())

	w.WriteHeader(200)
	w.Write([]byte("oi"))
}

func (h *httpHandlers) putHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("putHandler")
	h.service.UpdateBook(req.Context())

	w.WriteHeader(200)
}

func (h *httpHandlers) deleteHandler(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("deleteHandler")
	h.service.DeleteBook(req.Context())

	w.WriteHeader(200)
}

func NewHandler(logger logging.Logger, service services.BookService) HTTPHandlers {
	return &httpHandlers{service, logger, otel.Tracer("http handler")}
}
