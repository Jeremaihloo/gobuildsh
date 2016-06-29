package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"
)

// BuildInfo BuildInfo
type BuildInfo struct {
	Name     string
	CreateAt time.Time
}

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage : go-build-sh <go-excute-name>")
		return
	}
	info := BuildInfo{}
	info.Name = os.Args[1]
	info.CreateAt = time.Now()
	tpSrc, err := ioutil.ReadFile("build-template.sh")
	if err != nil {
		fmt.Println(err.Error())
	}
	t, err := template.New("gobuild").Parse(string(tpSrc))
	if err != nil {
		fmt.Println(err.Error())
	}
	var bs []byte
	w := bytes.NewBuffer(bs)
	err = t.Execute(w, info)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("ok! you can use ./build.sh {all|linux|darwin|windows} to build your go application")
	err = ioutil.WriteFile("build.sh", w.Bytes(), 0700)
	if err != nil {
		fmt.Println(err.Error())
	}
}
