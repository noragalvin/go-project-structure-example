package helpers

import (
	"runtime"
	"strings"
)

// WhereAmI ..
func WhereAmI(depthList ...int) (string, string, int) {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return chopPath(file), runtime.FuncForPC(function).Name(), line
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
