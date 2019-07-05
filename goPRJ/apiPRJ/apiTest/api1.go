package main
//
//import (
//	"encoding/json"
//	"log"
//	"strings"
//
//	"github.com/micro/go-micro"
//	"github.com/micro/go-micro/errors"
//	"goPRJ/apiPRJ/apiTest/proto"
//	proto "goPRJ/apiPRJ/apiTest/proto"
//	"golang.org/x/net/context"
//)
//
//type Example1 struct{}
//
//type Foo1 struct{}
//
//// Example.Call is a method which will be served by http request /example/call
//// In the event we see /[service]/[method] the [service] is used as part of the method
//// E.g /example/call goes to go.micro.api.example Example.Call
//func (e *Example) Call1(ctx context.Context, req *api.Request, rsp *api.Response) error {
//	log.Print("Received Example.Call request")
//
//	// parse values from the get request
//	name, ok := req.Get["name"]
//
//	if !ok || len(name.Values) == 0 {
//		return errors.BadRequest("go.micro.api.example", "no content")
//	}
//
//	// set response status
//	rsp.StatusCode = 200
//
//	// respond with some json
//	b, _ := json.Marshal(map[string]string{
//		"message": "got your request " + strings.Join(name.Values, " "),
//	})
//
//	// set json body
//	rsp.Body = string(b)
//
//	return nil
//}
//
//// Foo.Bar is a method which will be served by http request /example/foo/bar
//// Because Foo is not the same as the service name it is mapped beyond /example/
//func (f *Foo) Bar1(ctx context.Context, req *api.Request, rsp *api.Response) error {
//	log.Print("Received Foo.Bar request")
//
//	// check method
//	log.Print(req)
//	log.Print(req.Method)
//	if req.Method != "POST" {
//		log.Print(111)
//		return errors.BadRequest("go.micro.api.example", "require post")
//	}
//
//	// let's make sure we get json
//	ct, ok := req.Header["Content-Type"]
//	if !ok || len(ct.Values) == 0 {
//		return errors.BadRequest("go.micro.api.example", "need content-type")
//	}
//
//	if ct.Values[0] != "application/json" {
//		return errors.BadRequest("go.micro.api.example", "expect application/json")
//	}
//
//	// parse body
//	var body map[string]interface{}
//	json.Unmarshal([]byte(req.Body), &body)
//	rsp.Body="Hi, i have got your request!"
//	// do something with parsed body
//
//	return nil
//}
//
//func main() {
//	service := micro.NewService(
//		micro.Name("go.micro.api.example1"),
//	)
//
//	service.Init()
//
//	// register example handler
//	proto.RegisterExampleHandler(service.Server(), new(Example))
//
//	// register foo handler
//	proto.RegisterFooHandler(service.Server(), new(Foo))
//
//	if err := service.Run(); err != nil {
//		log.Fatal(err)
//	}
//}
//
////
////curl -H 'Content-Type: application/json' \
////-d '{"service": "go.micro.api.example", "method": "Foo.Bar","request":{} }'
////\
////http://localhost:8080/rpc