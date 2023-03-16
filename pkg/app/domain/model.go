package domain

// Options Main Object for the functions
type Options struct {
	Method  string // POST || GET
	Url     string
	Body    *string
	Auth    *AuthMethod
	Headers map[string]string
	Params  map[string]string
}
