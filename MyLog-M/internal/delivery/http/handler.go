package http

import (
	"MyLog-M/internal/domain"
	"MyLog-M/internal/repository"
	"MyLog-M/pkg/response"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//go:generate mockgen -source=./handler.go -destination=./mock/mock.go -package=mock
type service interface {
	Tail(limit int64) (*[]domain.Data, error)
}

type Handler struct {
	service service
}

//New Handler
func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

//MyTail returns as json the last limit number of records, with default limit=1
func (h *Handler) MyTail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		limit := 1 //default show lastest 1
		if n, err := strconv.Atoi(strings.ToUpper(r.URL.Query().Get("limit"))); err == nil {
			limit = n
			log.Println("limit:", limit)
		}
		resp, err := h.service.Tail(int64(limit))
		if err != nil {
			response.AsJSONError(w, 200, err.Error())
			return
		}
		log.Println(resp)
		response.AsJSON(w, 200, resp)
		return
	}
	response.AsJSONError(w, http.StatusMethodNotAllowed, "Invalid action")
}

var tplDir string = "./html/templates"

//LoadTemplate loads the template tmplName from tplDir
func LoadTemplate(tplDir, tmplName string) *template.Template {
	t, err := template.New(tmplName).ParseFiles(tplDir + "/" + tmplName)
	if err != nil {
		panic(err)
	}
	return t
}

//ViewData is a collection of data for the view
type ViewData struct {
	PageTitle string
	Records   []repository.Data
	RowCount  int
}
