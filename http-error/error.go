package httperror

import "net/http"

type ErrHttp int

func (e ErrHttp) Error() string {
	return http.StatusText(int(e))
}

func New(code int) ErrHttp {
	return ErrHttp(code)
}
