package helper

type ErrorValidation struct {
	Field        string
	ErrorMessage string
}

func ErrorMessageFromTag(input string) string {
	switch input {
	case "required":
		return "this field is required"
	case "email":
		return "Please provide a valid email address"
	case "min":
		return "Password should be at least 12 character long"
	}

	return ""
}
