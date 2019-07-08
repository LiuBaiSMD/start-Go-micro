package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/web"
	"io"
	"log"
	"net/http"
	"os"
	"html/template"
	"github.com/micro/go-micro/errors"
)

// exampleCall will handle /example/call
func exampleCall(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// get name
	name := r.Form.Get("name")

	if len(name) == 0 {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "no content").Error(),
			400,
		)
		return
	}

	// marshal response
	b, _ := json.Marshal(map[string]interface{}{
		"message": "got your message " + name,
	})

	// write response
	w.Write(b)
}

// exampleFooBar will handle /example/foo/bar
func exampleFooBar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "require post").Error(),
			400,
		)
		return
	}

	if len(r.Header.Get("Content-Type")) == 0 {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "need content-type").Error(),
			400,
		)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "expect application/json").Error(),
			400,
		)
		return
	}

	// do something
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t, _ := template.New("foo").Parse(`<html>
<head>
       <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://127.0.0.1:8080/example/foo/upload" method="post">
    <input type="file" name="uploadfile" />
   <br />
   保存目录： <input type="text" name="path" /> 如 /Users/me/Downloads/test/
     <br />
    <input type="submit" name='上传' value="upload" />
</form>
</body>
</html>`)
		t.Execute(w, nil)

		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	path := r.PostForm.Get("path")
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	// we're using go-web for convenience since it registers with discovery
	service := web.NewService(
		web.Name("go.micro.api.example"),
	)

	service.HandleFunc("/example/call", exampleCall)
	service.HandleFunc("/example/foo/bar", exampleFooBar)
	service.HandleFunc("/example/foo/upload", uploadFile)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
