package controllers

import (
	"net/http"
	//"log"
	"fmt"
	//"strings"
	"github.com/gorilla/mux"
	"path/filepath"
	"io/ioutil"
)

var page_cache map[string][]byte

func staticInit(){
	page_cache = make(map[string][]byte)
}

func static_controller(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	file := vars["file"]
	path, err := filepath.Abs("./assets/" + file)
	if err != nil {
		http.NotFound(w, r)
	}

	var byteFile []byte

	if page_cache[path] != nil{
		byteFile = page_cache[path]
	}else{
		byteFile, err = ioutil.ReadFile(path)
		if err != nil{
			http.NotFound(w, r)
		}

		page_cache[path] = byteFile
	}


	fmt.Fprint(w, string(byteFile))
}
