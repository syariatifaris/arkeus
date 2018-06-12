package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"strings"

	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/framework/entity"
	"github.com/syariatifaris/arkeus/core/framework/header"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/arkeus/core/retry"

	"github.com/syariatifaris/arkeus/core/net"
	"github.com/syariatifaris/arkeus/core/panics"
)

var retryPolicy = retry.NewRetryPolicy(MaxAttempt)

const MaxAttempt = 5

//Base Handler
type THandler interface {
	Name() string
	RegisterHandlers(router net.Router)
}

//The Application Mux Server
type MuxServer interface {
	//Http Handler
	http.Handler
	//Http Handler Func
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

//NewSimpleBaseHandler creates new  base handler
func NewSimpleBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

//Base Handler Implementation
//Put Dependencies Here
type BaseHandler struct {
	//Template Purpose
}

//HandlerResult Structure
type HandlerResult struct {
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage interface{} `json:"error,omitempty"`
	Status       int         `json:"status"`
}

// this is the func specification for all handle func of API-s that don't need to be validated,
// handle func don't need writer since writing/rendering will be handled by NoAuthenticate,
// handle func could easily just return data & error,
// we often forget to return after calling c.Renderxxx in error handling
type JsonHandlerFunc func(*http.Request) (data interface{}, err error)

//Region entry point handlers wrapper

//NoAuthenticate this is the wrapper for all handle func of API-s that don't need to be authenticated
func (b *BaseHandler) NoAuthenticate(f JsonHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer panics.Restore()
		data, err := f(r)
		b.renderJSON(w, data, err)
	}
}

//renderJSON renders in json format
func (b *BaseHandler) renderJSON(w http.ResponseWriter, data interface{}, err error) *entity.RenderResult {
	if err != nil {
		arklog.ERROR.Println("[BaseHandler][renderJson] Response Error:", err.Error())
	}

	r := constructHandlerResult(data, err)
	d, err := json.Marshal(r)
	if err != nil {
		return b.RenderError(w, err)
	}

	rr := &entity.RenderResult{
		HttpContentType: header.HttpContentTypeJSON,
		StatusCode:      r.Status,
		Body:            string(d),
	}

	b.writeResponse(w, rr)
	return rr
}

func (b *BaseHandler) writeResponse(w http.ResponseWriter, rr *entity.RenderResult) {
	w.Header().Set(header.HttpContentType, rr.HttpContentType)
	w.WriteHeader(rr.StatusCode)
	w.Write([]byte(rr.Body))
}

func constructHandlerResult(data interface{}, err error) HandlerResult {
	hr := HandlerResult{
		Status: http.StatusOK,
	}
	if data != nil {
		hr.Data = data
	}
	if err != nil {
		errType, httpStatus := errors.ErrorAndHTTPCode(err)
		hr.Status = httpStatus
		hr.ErrorMessage = errType
	}
	return hr
}

//RenderError renders the error with default code 500
func (b *BaseHandler) RenderError(w http.ResponseWriter, err error) *entity.RenderResult {
	if err != nil {
		message, code := errors.ErrorAndHTTPCode(err)
		arklog.ERROR.Printf("[BaseHandler][RenderError] %s. %s\n", http.StatusText(code), err.Error())
		rr := &entity.RenderResult{
			HttpContentType: header.HttpContentTypeHTML,
			StatusCode:      code,
			Body:            message,
		}
		b.writeResponse(w, rr)
		return rr
	}
	return nil
}

//RenderHTML renders HTML template
func (b *BaseHandler) RenderHTML(w http.ResponseWriter, templatePath string, data interface{}) {
	t := template.New("result template")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		arklog.ERROR.Printf("[BaseHandler][RenderHTML] Cannot parse the template. Err: %s\n", err.Error())
		b.RenderError(w, errors.New(errors.InternalServerError))
	}

	t.Execute(w, data)
}

//get post data, put data to an interface
func (b *BaseHandler) GetPostData(r *http.Request, v interface{}) error {
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("method %s not allowed %d", r.Method, http.StatusMethodNotAllowed)
	}

	return nil
}

//GroupOf check if
func GroupOf(name, prefix string) bool {
	return strings.Contains(name, prefix)
}

//RejectionHandler creates an action when request is rejected due exceeding limit
func (b *BaseHandler) RejectionHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		b.RenderError(writer, errors.New(errors.TooManyRequest))
	}
}
