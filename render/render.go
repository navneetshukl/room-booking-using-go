package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, url string){
	tmpl,_:=template.ParseFiles("./templates/"+url,"./templates/base.layout.tmpl")

	err:=tmpl.Execute(w,nil)
	if err!=nil{
		log.Fatal("Error in parsing template")
	}
}