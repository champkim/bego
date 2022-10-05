package app

import (
	"bego/conf"
	"bego/data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type AppHandler struct {
	//handler http.Handler
	http.Handler 
	db data.DBHandler
}

type Success struct {
	Success bool `json:"success"`	
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect) //인덱스 경로로 들어와도 todo.html 리다이렉션 해라 
}

func (a *AppHandler) getPagesHandler(w http.ResponseWriter, r *http.Request) {	
	list := a.db.GetPages()
	rd.JSON(w, http.StatusOK, list)
}

// func (a *AppHandler) addPageHandler(w http.ResponseWriter, r *http.Request) {
// 	page := new(data.Page)		
// 	err := json.NewDecoder(r.Body).Decode(page)
// 	if err != nil {		
// 		//rd.Text(w, http.StatusBadRequest, err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, err)	
// 		return
// 	}		
// 	//log.Println("test  >>> " + strconv.Itoa( page.Index) + " " + page.Contents)		
// 	a.db.AddPage(page)	
// 	rd.JSON(w, http.StatusCreated, page)
// }

// func (a *AppHandler) UpdatePageHandler(w http.ResponseWriter, r *http.Request) {
// 	page := new(data.Page)		
// 	err := json.NewDecoder(r.Body).Decode(page)
// 	if err != nil {		
// 		//rd.Text(w, http.StatusBadRequest, err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, err)	
// 		return
// 	}		

// 	ok := a.db.UpdatePage(page)

// 	if ok {
// 		rd.JSON(w, http.StatusOK, Success{true}) 
// 	} else {
// 		rd.JSON(w, http.StatusOK, Success{false})
// 	}
// }

func (a *AppHandler) addPageHandler(w http.ResponseWriter, r *http.Request) {	
	pages := []*data.Page{}
	err := json.NewDecoder(r.Body).Decode(&pages)
	if err != nil {		
		//rd.Text(w, http.StatusBadRequest, err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)	
		return
	}		
	//log.Println("test  >>> " + strconv.Itoa( page.Index) + " " + page.Contents)		
	var errNum int = 0 
	for _, p := range pages {		
		if !a.db.AddPage(p)	{
			errNum++
		}
	}

	if errNum == 0 {
		rd.JSON(w, http.StatusCreated, pages) 
	} else {
		rd.JSON(w, http.StatusBadRequest, nil)
	}
	//rd.JSON(w, http.StatusCreated, pages)
}

func (a *AppHandler) UpdatePageHandler(w http.ResponseWriter, r *http.Request) {
	pages := []*data.Page{}
	err := json.NewDecoder(r.Body).Decode(&pages)
	if err != nil {		
		//rd.Text(w, http.StatusBadRequest, err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)	
		return
	}		
	
	var errNum int = 0 
	for _, p := range pages {		
		ok := a.db.UpdatePage(p)
		if !ok {
			errNum++
		} 	
	}

	if errNum == 0 {
		rd.JSON(w, http.StatusOK, Success{true}) 
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}



func (a *AppHandler) GetPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["id"])
	//ok := model.RemoveTodo(id)
	//log.Println("index: " + strconv.Itoa(index))
	page := a.db.GetPage(index)

	if page != nil {
		rd.JSON(w, http.StatusOK, page) 
	} else {
		rd.JSON(w, http.StatusBadRequest, nil)
	}
}

func (a *AppHandler) DeletePageHandler(w http.ResponseWriter, r *http.Request) {
	ok := a.db.DeletePage()
	if ok  {		
		rd.JSON(w, http.StatusOK, nil) 
	} else {
		rd.JSON(w, http.StatusBadRequest, nil)		
	}
}

func (a *AppHandler) Close() {
	a.db.Close()	
}

// func corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Header().Set("Access-Control-Allow-Origin", "*")                                                            			       
// 			w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") 			       
// 			w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    
// 			w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		    w.Header().Add("Access-Control-Expose-Headers", "Origin,  X-Auth-Token, Authorization, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")	  			                                                
// 			w.Header().Set("content-type", "application/json;charset=UTF-8")
									
// 			if r.Method == "OPTIONS" {
// 					w.WriteHeader(http.StatusNoContent)
// 					return
// 			}
// 			next.ServeHTTP(w, r)
// 	})
// }

//func MakeHandler() http.Handler {
func MakeHandler(sconf *conf.ServerConf) *AppHandler {	

	mux := mux.NewRouter()
	//mux.Use(corsMiddleware)

	c := cors.New(cors.Options{
		AllowedHeaders:[]string{"*"}, 		
		AllowedOrigins:[]string{"*"}, 		
		AllowCredentials: false,		
		AllowedMethods: []string{"GET", "DELETE", "POST", "PUT", "OPTIONS"},		
	})

	a := &AppHandler{
		Handler: c.Handler(mux),
		db: data.NewDBHandler(sconf),		
	}
	
	mux.HandleFunc("/pages", a.getPagesHandler).Methods("GET")
	mux.HandleFunc("/pages", a.addPageHandler).Methods("POST")
	mux.HandleFunc("/pages", a.UpdatePageHandler).Methods("PUT")
	mux.HandleFunc("/pages", a.DeletePageHandler).Methods("DELETE")
	mux.HandleFunc("/pages/{id:[0-9]+}", a.GetPageHandler).Methods("GET")	
	mux.HandleFunc("/", a.indexHandler)
					
	return a
}