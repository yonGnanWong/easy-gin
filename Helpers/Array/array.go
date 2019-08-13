package Array

func InArray(s interface{}, d map[interface{}]interface{}) bool {
	for _, v := range d {
		if s == v {
			return true
		}
	}
	return false
}

func ArrayValues(d map[string]string) []string {
	slice := make([]string,len(d))

	for _, v := range d {
		if v != "" {
			slice = append(slice,v)
		}
	}
	return slice
}

//判断key是否存在map中
func MapKeyExist(needle string,d map[string]string) bool  {
	for k := range d {
		if k == needle {
			return true
		}
	}
	return false
}
