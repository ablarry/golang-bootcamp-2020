package controller

// ParamsURL is a struct with properties required in URL
type ParamsURL struct {
	ID int64 `form:":id"`
}
