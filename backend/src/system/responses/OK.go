package response

// Default response when a operation is success.
func OK() *Response {
	return &Response{Message: "Successful operation", Code: 200}
}

func OK_WITH_DATA(data interface{}) *Response {
	return &Response{Message: "Successful operation", Code: 200, Data: data}
}
