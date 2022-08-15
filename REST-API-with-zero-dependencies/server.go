package main
import (
	"net/http"
	"encoding/json"
)
type Coaster struct{
	Name string 	`json:"name"`
	Manufacturer string	`json:"manufacturer"`
	Id string 			`json:"id"`
	InPark string 		`json:"inpark"`
	Heigth int			`json:"heigth"`
}
type coasterHandlers struct{
	store map[string]Coaster
}
func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		h.get(w,r)
		return
	case "POST":
		h.post(w,r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request){

}
func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request){
	coasters:=make([]Coaster,len(h.store))
	i:=0
	for _, coaster:=range h.store{
		coasters[i]=coaster
		i++
	}
	jsonBytes, err:=json.Marshal(coasters)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type","application/json")
	w.Write(jsonBytes)

}
func newCoasterHandlers()*coasterHandlers{
	return &coasterHandlers{
		store:map[string]Coaster{
			"id1":Coaster{
				Name:"Fury 325",
				Heigth:99,
				Id:"id1",
				InPark:"Carowinds",
				Manufacturer:"B+M",
			},
		},
	}
}
func main(){
	coasterHandlers:=newCoasterHandlers()
	http.HandleFunc("/coasters",coasterHandlers.coasters)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		panic(err)
	}
}