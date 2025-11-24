package common

// HTTP Status Code Constants
const (
	// Success codes
	StatusOK      = 200
	StatusCreated = 201

	// Client error codes
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusPaymentRequired     = 402
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusRequestTimeout      = 408
	StatusConflict            = 409
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429

	// Server error codes
	StatusInternalServerError = 500
	StatusNotImplemented      = 501
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
	StatusGatewayTimeout      = 504
)

// API Operation Results
const (
	StatusSuccess = "success"
	StatusFailed  = "failed"
)
