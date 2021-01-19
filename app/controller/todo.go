package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"goServer/app"
	"io/ioutil"

	// "goTestServer/GoTestServer"
	"net/http"
)

type ToDoController struct {
	usecase app.ToDoUsecaseInterface
}

func setJsonResponse(w http.ResponseWriter, r *http.Request, httpCode int, resp interface{}) {
	w.WriteHeader(httpCode)
	r.Header.Set("Content-Type", "application/json")
	jsonB, _ := json.MarshalIndent(resp, "", " ")
	w.Write(jsonB)
}

func (td *ToDoController) CreateToDo(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			setJsonResponse(w, req, http.StatusInternalServerError, map[string]interface{}{"error": "internal server error"})
		}
	}()
	if req.Method == http.MethodPost {
		var t map[string]interface{}
		reqBody, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqBody, &t)
		ctx := req.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		fmt.Println(t, string(reqBody))
		resp, err := td.usecase.CreateToDo(ctx, t["title"].(string), t["description"].(string), t["priority"].(string), t["due_date"].(string), t["tag"].(string))
		if err != nil {
			setJsonResponse(w, req, http.StatusBadRequest, resp)
		} else {
			setJsonResponse(w, req, http.StatusOK, resp)
		}
	} else {
		setJsonResponse(w, req, http.StatusBadRequest, map[string]interface{}{"error": "method not allowed"})
	}
}

func (td *ToDoController) UpdateToDoById(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			setJsonResponse(w, req, http.StatusInternalServerError, map[string]interface{}{"error": "internal server error"})
		}
	}()
	if req.Method == http.MethodPut {
		fmt.Println(req.Method)
		var t map[string]interface{}
		ctx := req.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		reqBody, _ := ioutil.ReadAll(req.Body)
		err := json.Unmarshal(reqBody, &t)
		if err != nil {
			fmt.Println(err)
		}
		resp, err := td.usecase.UpdateToDoById(ctx, t["id"].(string), t["title"].(string), t["description"].(string), t["priority"].(string), t["due_date"].(string), t["is_completed"].(string), t["tag"].(string))
		if err != nil {
			setJsonResponse(w, req, http.StatusBadRequest, resp)
		} else {
			setJsonResponse(w, req, http.StatusOK, resp)
		}
	} else {
		setJsonResponse(w, req, http.StatusBadRequest, map[string]interface{}{"error": "method not allowed"})
	}
}

func (td *ToDoController) GetToDos(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			setJsonResponse(w, req, http.StatusInternalServerError, map[string]interface{}{"error": "internal server error"})
		}
	}()
	if req.Method == http.MethodGet {

		limit := req.URL.Query().Get("limit")
		offset := req.URL.Query().Get("offset")
		ctx := req.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		resp, err := td.usecase.GetToDos(ctx, limit, offset)
		if err != nil {
			setJsonResponse(w, req, http.StatusBadRequest, resp)
		} else {
			setJsonResponse(w, req, http.StatusOK, resp)
		}
	} else {
		setJsonResponse(w, req, http.StatusBadRequest, map[string]interface{}{"error": "method not allowed"})
	}
}
func (td *ToDoController) DeleteToDoById(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			setJsonResponse(w, req, http.StatusInternalServerError, map[string]interface{}{"error": "internal server error"})
		}
	}()
	if req.Method == http.MethodDelete {
		ctx := req.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		id := req.URL.Query().Get("id")
		resp, err := td.usecase.DeleteToDoById(ctx, id)
		if err != nil {
			setJsonResponse(w, req, http.StatusBadRequest, resp)
		} else {
			setJsonResponse(w, req, http.StatusOK, resp)
		}
	} else {
		setJsonResponse(w, req, http.StatusBadRequest, map[string]interface{}{"error": "method not allowed"})
	}
}

func NewTodo(s *http.ServeMux, u app.ToDoUsecaseInterface) {
	hand := &ToDoController{
		usecase: u,
	}
	s.HandleFunc("/create-todo", hand.CreateToDo)
	s.HandleFunc("/update-todo-by-id", hand.UpdateToDoById)
	s.HandleFunc("/get-todos", hand.GetToDos)
	s.HandleFunc("/delete-todo-by-id", hand.DeleteToDoById)
}
