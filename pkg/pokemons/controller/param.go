package controller

type ParamsURL struct {
	Id int64 `form:":id" validate:"required,gt=0"`
}