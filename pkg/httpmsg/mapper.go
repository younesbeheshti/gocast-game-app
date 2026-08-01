package httpmsg

import (
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"net/http"
)

func Error(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		var msg string
		code := mapKindToHTTPStatusCode(re.Kind())
		if code >= 500 {
			msg = "internal server error"
		} else {
			msg = re.Message()
		}
		return msg, code

	default:
		return err.Error(), http.StatusBadRequest
	}

}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}

}
