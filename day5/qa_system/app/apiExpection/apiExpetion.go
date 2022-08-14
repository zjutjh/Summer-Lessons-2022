package apiExpection

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	ServerError  = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	ParamError   = NewError(http.StatusInternalServerError, 200501, "参数错误")
	HttpTimeout  = NewError(http.StatusInternalServerError, 200505, "系统异常，请稍后重试!")
	RequestError = NewError(http.StatusInternalServerError, 200506, "系统异常，请稍后重试!")
	ReSubmit     = NewError(http.StatusOK, 200507, "重复提交")
	TimeOut      = NewError(http.StatusOK, 200508, "超出试卷可作答时间")
	NotFound     = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	Unknown      = NewError(http.StatusInternalServerError, 300500, "系统异常，请稍后重试!")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}
