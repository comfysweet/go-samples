package main

import (
	"encoding/json"
	"fmt"
	"github.com/comfysweet/go-samples/day8/model"
	"log"
	"net/http"
	"strings"
)

type server struct {
}

func getName(req *http.Request) string {
	path := strings.Split(req.URL.Path, "/")
	return path[len(path)-1]
}

type contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	defer model.CloseDB()
	http.Handle("/phones/", &server{})
	http.ListenAndServe(":8080", nil)
}

func (s *server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "POST":
		var c contact
		json.NewDecoder(req.Body).Decode(&c)
		fmt.Fprintf(resp, "name: %v, phone: %v\n", c.Name, c.Phone)
		id := model.Create(c.Name, c.Phone)
		resp.WriteHeader(http.StatusCreated)
		fmt.Fprintf(resp, `{"id":%d}`, id)

	case "GET":
		name := getName(req)
		if name == "" {
			contacts := model.ReadAll()
			res := make([]contact, len(contacts))
			for i, c := range contacts {
				res[i] = contact{c[0], c[1]}
			}
			jsonResult, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			resp.WriteHeader(http.StatusOK)
			fmt.Fprint(resp, string(jsonResult))
		} else {
			phone := model.Read(name)
			if phone != "" {
				jsonResult, err := json.Marshal(&contact{name, phone})
				if err != nil {
					log.Fatal(err)
				}
				resp.WriteHeader(http.StatusOK)
				fmt.Fprint(resp, string(jsonResult))
			} else {
				resp.WriteHeader(http.StatusNotFound)
			}
		}
	}
}
