package main

import (
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var (
	server    *gin.Engine
	templates *template.Template
	layouts   map[string]string
	includes  map[string]string
	buffer    []byte
)

func createMyRender() (renderer multitemplate.Renderer) {
	renderer = multitemplate.NewRenderer()
	layouts = make(map[string]string)
	includes = make(map[string]string)

	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		if strings.HasPrefix(name, "/layouts/") {
			buffer, err = ioutil.ReadAll(file)
			if err != nil {
				return nil
			}
			layouts["name"] = string(buffer)
		}
		if strings.HasPrefix(name, "/includes/") {
			buffer, err = ioutil.ReadAll(file)
			if err != nil {
				return nil
			}
			includes[name] = string(buffer)
		}
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for name, include := range includes {
		name = name[len("/includes/"):]
		tmpl := template.New(name)
		_, err = tmpl.Parse(include)
		for _, layout := range layouts {
			_, err = tmpl.Parse(layout)
		}
		renderer.Add(name, tmpl)
	}
	return
}
