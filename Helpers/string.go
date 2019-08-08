package Helpers

func IsString(p interface{}) bool {
	switch p.(type) {
	case string:
		return true
	default:
		return false
	}
}
