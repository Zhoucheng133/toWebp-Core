package main

import "towebp_core/utils"
import "C"

//export Convert
func Convert(path *C.char, width, height *C.int, output *C.char) *C.char {
	err := utils.Convert(C.GoString(path), int(C.int(*width)), int(C.int(*height)), C.GoString(output))
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString("ok")
}

func main() {}
