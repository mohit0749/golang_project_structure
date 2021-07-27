package util

func SliceInterfaceToString(args ...interface{}) string {
	str := ""
	for _, arg := range args {
		msg, _ := arg.(string)
		str += " " + msg
	}
	return str
}
