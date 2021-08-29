package http

import (
	"MyLog-M/internal/repository"
	"MyLog-M/pkg/response"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	db *sql.DB
}

//New Handler
func New(db *sql.DB) *Handler {
	return &Handler{db: db}

}

//MyLog gets a record by id, or insert a posted record
func (h Handler) MyLog(w http.ResponseWriter, r *http.Request) {
	store := repository.NewLogRepo(h.db)
	if r.Method == http.MethodGet {
		ID := -1
		if n, err := strconv.Atoi(strings.ToUpper(r.URL.Query().Get("id"))); err == nil {
			ID = n
		}
		if ID < 0 {
			response.AsJSONError(w, http.StatusMethodNotAllowed, "Invalid id")
			return
		}
		resp, err := store.Get(int64(ID))
		log.Println(resp)
		if err != nil {
			response.AsJSONError(w, http.StatusMethodNotAllowed, err.Error())
			return
		}
		response.AsJSON(w, 200, resp)
		return
	}
	if r.Method == http.MethodPost {
		var data repository.Data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			response.AsJSONError(w, 200, err.Error())
			return
		}
		LastID, err := store.Insert(data)
		if err != nil {
			response.AsJSONError(w, 200, err.Error())
			return
		}
		resp, err := store.Get(LastID)
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

//MyTail returns as json the last limit number of records, with default limit=1
func (h Handler) MyTail(w http.ResponseWriter, r *http.Request) {
	store := repository.NewLogRepo(h.db)
	if r.Method == http.MethodGet {
		limit := 1 //default show lastest 1
		if n, err := strconv.Atoi(strings.ToUpper(r.URL.Query().Get("limit"))); err == nil {
			limit = n
			log.Println("limit:", limit)
		}
		resp, err := store.Tail(int64(limit))
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

//MyView shows the last limit number of records, with default limit=1
func (h Handler) MyView(w http.ResponseWriter, r *http.Request) {
	tmpl := LoadTemplate(tplDir, "view.gohtml")
	store := repository.NewLogRepo(h.db)
	if r.Method == http.MethodGet {
		limit := 1 //default show lastest 1
		if n, err := strconv.Atoi(strings.ToUpper(r.URL.Query().Get("limit"))); err == nil {
			limit = n
		}
		recs, err := store.Tail(int64(limit))
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		viewData := &ViewData{PageTitle: "VIEW", Records: *recs, RowCount: len(*recs)}
		tmpl.Execute(w, viewData)
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
