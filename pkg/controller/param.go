package controller

// ParamsOne is a struct with properties required in URL
type ParamsOne struct {
	ID int64 `form:":id" validate:"required,gt=0"`
}

// ParamsList is a struct with properties required in URL
type ParamsList struct {
	Size      int64 `form:"size"`
	Paginator int64 `form:"paginator"`
}
