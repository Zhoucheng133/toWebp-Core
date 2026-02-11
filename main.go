package main

import (
	"encoding/json"
	"fmt"
	"towebp_core/utils"
)
import "C"

//export Convert
func Convert(path *C.char, width, height *C.int, output *C.char) *C.char {
	err := utils.Convert(C.GoString(path), int(C.int(*width)), int(C.int(*height)), C.GoString(output))
	if err != nil {
		return C.CString(fmt.Sprint("ERR: ", err.Error()))
	}
	return C.CString("OK")
}

//export Scan
func Scan(path *C.char) *C.char {
	jsonData, err := json.Marshal(utils.Scan(C.GoString(path)))
	if err != nil {
		return C.CString(fmt.Sprint("ERR: ", err.Error()))
	}
	return C.CString(string(jsonData))
}

//export ConvertFromDir
func ConvertFromDir(path *C.char, width, height *C.int, output *C.char) *C.char {
	err := utils.ConvertFromDir(C.GoString(path), int(C.int(*width)), int(C.int(*height)), C.GoString(output))
	if err != nil {
		return C.CString(fmt.Sprint("ERR: ", err.Error()))
	}
	return C.CString("OK")
}

func main() {
	fmt.Println(utils.Scan("/Users/zhoucheng/Downloads/照片"))
}
