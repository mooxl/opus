package model

import (
	"io/ioutil"
)

// GetTemplates :
func GetTemplates(file string) []string {
	templates := []string{file}
	files, _ := ioutil.ReadDir("./view/templates")
	for _, template := range files {
		templates = append(templates, "./view/templates/"+template.Name())
	}
	return templates
}
