package router

/*import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"net/http"
)

var (
	v       = validator.New(validator.WithRequiredStructEnabled())
	decoder = schema.NewDecoder()
)

func decodeAndValidate[reqType any](r *http.Request, req reqType) (err error) {
	if err = json.NewDecoder(r.Body).Decode(req); err != nil {
		return
	}
	return v.Struct(req)
}

func DecodeQueryAndValidate[reqType any](r *http.Request, req reqType) (err error) {
	if err = decoder.Decode(req, r.URL.Query()); err != nil {
		return
	}
	return v.Struct(req)
}
*/
