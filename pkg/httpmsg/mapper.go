package httpmsg

import (
	"errors"
	"github.com/mobin-alz/gameapp/pkg/errmsg"
	"github.com/mobin-alz/gameapp/pkg/richerror"
	"net/http"
)

func Error(err error) (message string, code int) {
	var richError richerror.RichError
	switch {
	case errors.As(err, &richError):
		var re richerror.RichError
		errors.As(err, &re)
		code := mapKindToHTTPStatusCode(re.Kind())
		msg := re.Message()
		if code >= 500 {
			msg = errmsg.ErrorMsgSomethingWentWrong
		}
		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
