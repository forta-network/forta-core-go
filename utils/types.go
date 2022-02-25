package utils

// String converts string ptr to string.
func String(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}

// StringPtr converts string to string ptr.
func StringPtr(str string) *string {
	return &str
}

// Int32Ptr returns a pointer.
func Int32Ptr(n int32) *int32 {
	return &n
}

// BoolPtr returns a pointer.
func BoolPtr(v bool) *bool {
	return &v
}

// Bool returns a value.
func Bool(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}
