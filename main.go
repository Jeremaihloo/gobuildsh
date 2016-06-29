package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"
)

var tpSrc = `###################Start Script#################  
#!/bin/bash  
# build.sh to build go application for multi platform
# created by http://github.com/jeremaihloo/go-build-sh at {{ .CreateAt }}

if [ ! -d "dist" ]; then
    mkdir dist
fi

windows(){
    # windows
    export GOARCH=386
    export GOOS=windows
    go build -o {{ .Name }}.exe
    if [ ! -d "dist/windows" ]; then
        mkdir dist/windows
    fi
    mv {{ .Name }}.exe dist/windows/
}

linux(){
    # linux
    export GOARCH=amd64
    export GOOS=linux
    go build -o {{ .Name }}
    if [ ! -d "dist/linux" ]; then
        mkdir dist/linux
    fi
    mv {{ .Name }} dist/linux/
}

darwin(){
    # darwin
    export GOARCH=amd64
    export GOOS=darwin
    go build -o {{ .Name }}
    if [ ! -d "dist/darwin" ]; then
        mkdir dist/darwin
    fi
    mv {{ .Name }} dist/darwin/
}

all(){
    windows
    linux
    darwin    
}

case "$1" in
all)
    all    
    ;;
windows)
    windows
    ;;
linux)
    linux
    ;;
darwin)
    darwin
    ;;
*)
    echo $"Usage: $0 {all|windows|darwin|linux}"
    exit 1
esac

exit 0
#####################End Script##################


`

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
