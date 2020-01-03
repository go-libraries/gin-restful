package gen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Build struct {
	Port          string
	Dsn           string
	ProjectName   string
	ProjectPath   string
	TemplatePath  string
	TemplateFiles map[string]string
}

func (build *Build) Load(TemplatePath string) {
	if TemplatePath != "" {
		build.TemplatePath = TemplatePath + "/templates"
	} else {
		build.TemplatePath, _ = os.Getwd()
		build.TemplatePath += "/templates"
	}

	path := build.TemplatePath
	build.load(path)
}

func (build *Build) load(path string) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {

			build.load(path +"/"+ file.Name() + "/")
		} else {
			path := strings.TrimRight(path, "/")
			build.makeTemplate(path +"/" + file.Name())
		}
	}
}

func (build *Build) makeTemplate(filename string) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		panic("can't load file :" + filename)
	}

	bt, err := ioutil.ReadAll(f)
	if err != nil {
		panic("can't load file :" + filename)
	}
	//strings.Replace(build.TemplatePath, "//", "/", -1)
	key := strings.Replace(filename, build.TemplatePath, "", -1)
	build.TemplateFiles[key] = string(bt[:])
}

func (build *Build) Generator() {
	for filename, content := range build.TemplateFiles {
		filename = build.ProjectPath + "/" + build.ProjectName + filename

		content := strings.Replace(content, "{{package}}", build.ProjectName, -1)
		content = strings.Replace(content, "{{dsn}}", build.Dsn, -1)
		content = strings.Replace(content, "{{port}}", build.Port, -1)
		filename = strings.Replace(filename, ".template", ".go", -1)
		writeFile(filename, content)
	}
}

func writeFile(filename, content string) {
	err := os.MkdirAll(getDir(filename), os.ModePerm)
	if err != nil {
		panic("can't create path:" + getDir(filename))
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		panic("can't create file:" + filename)
	}

	defer f.Close()
	_,err = f.WriteString(content)
	if err != nil {
		panic("can't write file:" + filename)
	}
}

func getDir(filename string) string {
	str, _ := filepath.Abs(filepath.Dir(filename))
	return str
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
