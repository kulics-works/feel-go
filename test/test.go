package webkit

import "reflect"
import "strings"
import "github.com/gin-gonic/gin"

type Context = *gin.Context
type WebServer struct {
	listen string
	engine *gin.Engine
}

func NewWebServerDefault(listen string) (v *WebServer) {
	srv := &WebServer{}
	srv.listen = listen
	srv.engine = gin.Default()
	return srv
}
func (me *WebServer) Run() (e error) {
	return me.engine.Run(me.listen)
}
func (me *WebServer) HandleFunc(method Method, relativePath string, handle func(Context)) (v *WebServer) {
	switch method {
	case GET:
		{
			me.engine.GET(relativePath, handle)
		}

	case POST:
		{
			me.engine.POST(relativePath, handle)
		}

	case PUT:
		{
			me.engine.PUT(relativePath, handle)
		}

	case DELETE:
		{
			me.engine.DELETE(relativePath, handle)
		}

	case PATCH:
		{
			me.engine.PATCH(relativePath, handle)
		}

	case OPTIONS:
		{
			me.engine.OPTIONS(relativePath, handle)
		}

	}
	return me
}
func (me *WebServer) HandleGET(relativePath string, handle func(Context)) (v *WebServer) {
	return me.HandleFunc(GET, relativePath, handle)
}
func (me *WebServer) HandlePOST(relativePath string, handle func(Context)) (v *WebServer) {
	return me.HandleFunc(POST, relativePath, handle)
}
func (me *WebServer) HandlePUT(relativePath string, handle func(Context)) (v *WebServer) {
	return me.HandleFunc(PUT, relativePath, handle)
}
func (me *WebServer) HandleDELETE(relativePath string, handle func(Context)) (v *WebServer) {
	return me.HandleFunc(DELETE, relativePath, handle)
}
func (me *WebServer) HandlePATCH(relativePath string, handle func(Context)) (v *WebServer) {
	return me.HandleFunc(PATCH, relativePath, handle)
}
func (me *WebServer) HandleOPTIONS(relativePath string, handle func(Context)) (v *WebServer) {
	return me.Handle_Func(OPTIONS, relativePath, handle)
}
func (me *WebServer) HandleStruct(relativePath string, handle interface{}) (v *WebServer) {
	me.handleStruct(relativePath, handle)
	return me
}
func (me *WebServer) handleStruct(relativePath string, handle interface{}) {
	rfType := reflect.TypeOf(handle)
	rfValue := reflect.ValueOf(handle)
	switch rfType.Kind() {
	case reflect.Ptr:
		{
			for i := 0; i < rfType.NumMethod(); i += 1 {
				methodName := rfType.Method(i).Name
				if !isMethod(NewMethod(methodName)) {
					continue
				}
				handleFunc := func(ctx Context) {
					rfValue.MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(ctx)})
				}
				me.HandleFunc(NewMethod(methodName), relativePath, handleFunc)
			}
			for i := 0; i < rfType.NumField(); i += 1 {
				fieldName := rfType.Field(i).Name
				me.handleStruct(relativePath+"/"+strings.ToLower(fieldName[:1])+fieldName[1:], rfValue.FieldByName(fieldName).Interface())
			}
		}
	case reflect.Interface:
		{
			for i := 0; i < rfType.NumMethod(); i += 1 {
				methodName := rfType.Method(i).Name
				if !isMethod(NewMethod(methodName)) {
					continue
				}
				handleFunc := func(ctx Context) {
					rfValue.MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(ctx)})
				}
				me.HandleFunc(NewMethod(methodName), relativePath, handleFunc)
			}
			for i := 0; i < rfType.NumField(); i += 1 {
				fieldName := rfType.Field(i).Name
				me.handleStruct(relativePath+"/"+strings.ToLower(fieldName[:1])+fieldName[1:], rfValue.FieldByName(fieldName).Interface())
			}
		}
	case reflect.Struct:
		{
			for i := 0; i < rfType.NumMethod(); i += 1 {
				methodName := rfType.Method(i).Name
				if !isMethod(NewMethod(methodName)) {
					continue
				}
				handleFunc := func(ctx Context) {
					rfValue.MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(ctx)})
				}
				me.HandleFunc(NewMethod(methodName), relativePath, handleFunc)
			}
			for i := 0; i < rfType.NumField(); i += 1 {
				fieldName := rfType.Field(i).Name
				me.handleStruct(relativePath+"/"+strings.ToLower(fieldName[:1])+fieldName[1:], rfValue.FieldByName(fieldName).Interface())
			}
		}

	}
}
