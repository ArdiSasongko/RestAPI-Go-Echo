package helper

func DefaultEmpty(value, defaultValue interface{}) interface{} {
	switch v := value.(type) {
	case string:
		if v == "" {
			return defaultValue
		}
		return v
	case int:
		if v == 0 {
			return defaultValue
		}
		return v
	default:
		return defaultValue
	}
}
