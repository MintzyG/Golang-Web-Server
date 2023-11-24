package main 

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"html/template"
)

type inscritos struct{
  ID int
  Nome string
  Cursos int
}

type pageData struct{
  Usuarios []inscritos
}

var Pessoas = []inscritos{
  {ID:01, Nome:"Zadoque", Cursos:2},
}

var templ *template.Template

func main(){
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.HandleFunc("/", home)
  http.HandleFunc("/dashboard/", dashboard)
  http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, _ *http.Request){
  templ = template.Must(template.ParseFiles("template/home.gohtml"))
  var data pageData
  data.Usuarios = append(data.Usuarios, Pessoas...)
  templ.Execute(w, data)
}

func dashboard(w http.ResponseWriter, r *http.Request){
  id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path,"/dashboard/"))
  decoder := json.NewDecoder(r.Body)
  switch met := r.Method; met {
  case "GET":
    w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(Pessoas)
  case "POST":
    var i inscritos
    decoder.Decode(&i)
    Pessoas = append(Pessoas,i)
  case "DELETE":
    var index int
    for i, p := range Pessoas{
      if p.ID == id {
        index = i
      } 
    }
    Pessoas = append(Pessoas[:index], Pessoas[index+1:]...)
  default:
    fmt.Printf("Código falhou miseravelmente")
  }
}
