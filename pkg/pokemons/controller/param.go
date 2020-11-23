package controller

// ParamsURL is a struct with properties required in URL
type ParamsURL struct {
	Id int64 `form:":id"`
}
