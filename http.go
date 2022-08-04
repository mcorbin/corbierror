package corbierror

var HTTPCodes = map[ErrorType]int{
	BadRequest:   400,
	Unauthorized: 401,
	Forbidden:    403,
	NotFound:     404,
	Conflict:     409,
	Internal:     500,
}
