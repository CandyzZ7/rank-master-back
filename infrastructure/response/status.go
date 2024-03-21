package response

var (
	OK          = New(0, "Success")
	ServerError = New(500, "Server Error")
)
