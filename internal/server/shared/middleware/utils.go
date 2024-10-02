package middleware

/*import (
	"bytes"
	"gitlab.com/piorun102/lg"
	"io"
	"net/http"
	"ranx_form/internal/server/shared"
	"ranx_form/pkg/http/codes"
	"strconv"
	"strings"
	"time"
)

func validateBody(r *http.Request, ctx lg.CtxLogger) (body []byte, code int) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Debugf("Error read body")
		return nil, codes.BodyRead
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return body, codes.NoError
}

func validateParam(r *http.Request) (string, int) {
	id := r.URL.Query().Get(shared.ID)
	if id == "" {
		lg.Errorf("Error id from params is ''")
	}
	return id, codes.NoError
}

func validateNumMonth(r *http.Request, ctx lg.CtxLogger) (int, int, int) {
	month, err := strconv.Atoi(r.URL.Query().Get(shared.NumMonth))
	if err != nil {
		ctx.Debugf("Error read query month params")
		return 0, 0, codes.ErrorReadQueryParam
	}

	year, err := strconv.Atoi(r.URL.Query().Get(shared.Year))
	if err != nil {
		ctx.Debugf("Error read query year params")
		return 0, 0, codes.ErrorReadQueryParam
	}

	code := validateYear(year)
	if code != codes.NoError {
		return 0, 0, codes.NotValidYear
	}

	return month, year, codes.NoError
}

func validateYear(year int) int {
	if year < 1000 || year > 10000 {
		return codes.NotValidYear
	} else if year > time.Now().Year() {
		return codes.NotValidYear
	}

	return codes.NoError
}

func validateCode(r *http.Request, ctx lg.CtxLogger) (string, int) {
	code := r.Header.Get(shared.Code)
	if code == "" {
		ctx.Debugf("Error read query code params")
		return "", codes.ErrorReadQueryParam
	}
	return code, codes.NoError
}

func validateAccessToken(r *http.Request, ctx lg.CtxLogger) (string, int) {
	accessTokenNotSplit := r.Header.Get(shared.AccessToken)

	accessToken := strings.Split(accessTokenNotSplit, " ")
	if len(accessToken) < 2 {
		ctx.Debugf("acess token not found")
		return "", codes.ErrorReadQueryParam
	}

	if accessToken[1] == "" {
		ctx.Debugf("Error read query accessToken params")
		return "", codes.ErrorReadQueryParam
	}

	return accessToken[1], codes.NoError
}

func validateRefreshToken(r *http.Request, ctx lg.CtxLogger) (string, int) {
	cookies := r.Cookies()
	if len(cookies) < 1 {
		ctx.Debugf("not found refresh_token from cookie")
		return "", codes.NotFoundRefreshToken
	}

	if cookies[0].Name != "refresh_token" {
		ctx.Debugf("error, title not match %s", cookies[0].Name)
		return "", codes.NotFoundTitleCookies
	}

	return cookies[0].Value, codes.NoError
}
*/
