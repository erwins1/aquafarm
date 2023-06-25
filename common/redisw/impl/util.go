package impl

func SliceOfStringsToSliceOfInterfaces(strs []string) []interface{} {
	var intf []interface{}

	for _, str := range strs {
		intf = append(intf, str)
	}

	return intf
}

func ConvertStringToInterface(value string) (result interface{}) {
	result = value
	return
}
