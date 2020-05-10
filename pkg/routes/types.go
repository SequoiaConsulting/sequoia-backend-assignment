package routes

// A response that indicates error in performed operation
// swagger:response errored
type errorResponse struct {
	// in: body
	// required: true
	Body struct {
		Error string `json:"error"`
	}
}

// Returned on successful resource creation. It also contains the path of created
// resource on the server.
// swagger:response created
type createdResponse struct {
	// in: body
	// required: false
	Body struct {
		Href string `json:"href"`
	}
}

// A list of requested resources with their location
// swagger:response listResponse
type listResponse struct {
	// in: body
	Body struct {
		Hrefs []string `json:"hrefs"`
	}
}

var (
	// ErrInvalidBody is a server response when it is not able to properly
	// decode the request body
	ErrInvalidBody = generateErrorResponse("unable to decode request body")

	// ErrInternalServerError is a server response when some unexpected error
	// is encoutered during operation
	ErrInternalServerError = generateErrorResponse("internal server error")
)
