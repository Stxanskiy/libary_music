package middleware

/*import (
	"context"
	"fmt"
	lb "gitlab.com/nevasik7/tglib"
	"gitlab.com/piorun102/lg"
	"net/http"
	i "ranx_form/internal/activity/interface"
	"ranx_form/internal/server/shared"
	"ranx_form/pkg/api/ranepa"
	"ranx_form/pkg/http/codes"
	"ranx_form/pkg/http/response"
	"strconv"
)

type I interface {
	VerifyCode(next http.Handler) http.Handler
	VerifyBody(next http.Handler) http.Handler
	VerifyParams(next http.Handler) http.Handler
	VerifyMonth(next http.Handler) http.Handler
	VerifyAccessToken(next http.Handler) http.Handler
	VerifyAdmin(next http.Handler) http.Handler
	VerifyRefreshToken(next http.Handler) http.Handler
}

type mw struct {
	curatorRepo  i.UserRepo
	activityRepo i.ActivityRepo
	ranepaUC     ranepa.UC
}

func New(curatorRepo i.UserRepo, activityRepo i.ActivityRepo, ranepaUC ranepa.UC) I {
	return &mw{
		curatorRepo:  curatorRepo,
		activityRepo: activityRepo,
		ranepaUC:     ranepaUC,
	}
}

func (m *mw) VerifyBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = lg.Ctx(r.Context(), r.Header)
			err error
		)
		defer ctx.DefError(&err)

		rCopy := r.Clone(r.Context())

		body, code := validateBody(rCopy, ctx)
		if code != codes.NoError {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка чтения body в middleware. Ошибка на стороне клиента. Code=%v", code), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error validate body from verifyBody")
			response.Write(w, codes.BodyRead, nil)
			return
		}

		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Method, r.Method))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Path, r.URL.Path))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Body, string(body)))

		next.ServeHTTP(w, rCopy)
	})
}

func (m *mw) VerifyParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = lg.Ctx(r.Context(), r.Header)
			err error
		)
		defer ctx.DefError(&err)

		rCopy := r.Clone(r.Context())

		curatorID, code := validateParam(rCopy)
		if code != codes.NoError {
			lg.Errorf("Error validate param ID from VerifyParams")
			if err = lb.SendRTg(fmt.Sprintf("Ошибка валидации параметра id в middleware. Ошибка из клиентской части. Code=%v", code), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}
			response.Write(w, codes.ErrorReadQueryParam, nil)
			return
		}

		idCurator, err := strconv.Atoi(curatorID)
		if err != nil {
			lg.Errorf("Error convert to num curatorID=%v", curatorID)
		}

		_, err = m.curatorRepo.FindCuratorById(ctx, int64(idCurator))
		if err != nil {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware, не найдено куратора по такому id=%v. Code=%v", idCurator, code), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error not found curator by curatorID=%v", idCurator)
			response.Write(w, codes.CuratorNotFound, nil)
			return
		}

		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.ID, idCurator))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Method, r.Method))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Path, r.URL.Path))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Body, ""))

		next.ServeHTTP(w, rCopy)
	})
}

func (m *mw) VerifyMonth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = lg.Ctx(r.Context(), r.Header)
			err error
		)
		defer ctx.DefError(&err)

		rCopy := r.Clone(r.Context())

		month, year, code := validateNumMonth(rCopy, ctx)
		if code != codes.NoError {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware по статистике. Не удалось считать месяц или год. Ошибка с клиентской стороны"), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}
			lg.Errorf("Error validate month and year from VerifyMonth")
			response.Write(w, codes.ErrorReadQueryParam, nil)
			return
		}

		if month < 1 || month > 12 {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware по статистике. Невалидный месяц: %v", month), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error number of a month outside < 1 or > 12 from VerifyMonth")
			response.Write(w, codes.UnknownMonth, nil)
			return
		}

		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Method, r.Method))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Path, r.URL.Path))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Body, ""))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.NumMonth, month))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Year, year))

		next.ServeHTTP(w, rCopy)
	})
}

func (m *mw) VerifyCode(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = lg.Ctx(r.Context(), r.Header)
			err error
		)
		defer ctx.DefError(&err)

		rCopy := r.Clone(r.Context())

		value, code := validateCode(rCopy, ctx)
		if code != codes.NoError {
			lg.Errorf("Error validate code from VerifyCode")
			if err = lb.SendRTg(fmt.Sprintf("Ошибка валидации кода в middleware на ручку /signIn из клиентской части. Code=%v", code), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			response.Write(w, codes.ErrorReadQueryParam, nil)
			return
		}

		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Method, r.Method))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Path, r.URL.Path))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Body, ""))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Code, value))

		next.ServeHTTP(w, rCopy)
	})
}

func (m *mw) VerifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = lg.Ctx(r.Context(), r.Header)
			err error
		)
		defer ctx.DefError(&err)

		rCopy := r.Clone(r.Context())

		accessToken, code := validateAccessToken(rCopy, ctx)
		if code != codes.NoError {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка валидации access token в middleware. Ошибка на стороне клиента. Code=%v", code), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error validate access token from VerifyAccessToken")
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		client, code := m.ranepaUC.GetClient(accessToken)
		if code != codes.NoError {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware. По такому access token=%v не удалось найти информацию по куратору в bitrixAPI", accessToken), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error GetClient to musiclibary api by access token=%v", accessToken)
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		if client.Result.Id == "" {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware. C bitrixAPI нашелся куратор, но с пустым id по такому accessToken=%v. Response=%v", accessToken, client.Result), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error clientID= '' by access token=%v", accessToken)
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		lg.Infof("clientID from response by access token=%v", accessToken)

		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Method, r.Method))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Path, r.URL.Path))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Body, ""))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.AccessToken, accessToken))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.ID, client.Result.Id))

		next.ServeHTTP(w, rCopy)
	})
}

func (m *mw) VerifyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = lg.Ctx(r.Context(), r.Header)
			err error
		)
		defer ctx.DefError(&err)

		rCopy := r.Clone(r.Context())

		accessToken, code := validateAccessToken(rCopy, ctx)
		if code != codes.NoError {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware при чтении access token на валидацию по админу. Ошибка на стороне клиента"), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error validate access token from VerifyAccessToken")
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		client, code := m.ranepaUC.GetClient(accessToken)
		if code != codes.NoError {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware. По такому access token=%v не удалось найти информацию по куратору в bitrixAPI", accessToken), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error GetClient to musiclibary api by access token=%v", accessToken)
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		if client.Result.Id == "" {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware. C bitrixAPI нашелся куратор, но с пустым id по такому accessToken=%v. Response=%v", accessToken, client.Result), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error clientID= '' by access token=%v", accessToken)
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		curatorID, err := strconv.Atoi(client.Result.Id)
		if err != nil {
			lg.Errorf("Error convert clientID from api adminMiddleware%v", client.Result.Id)
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		roleID, err := m.curatorRepo.FindCuratorById(ctx, int64(curatorID))
		if roleID != 1 {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware. Куратор с id=%v не является админом. Пожалуйста установите ему админа в бд или же игнорируйте эту ошибку", curatorID), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Curator not admin, id=%v", curatorID)
			response.Write(w, codes.Unauthorized, nil)
			return
		}

		lg.Infof("clientID from response by access token=%v", accessToken)

		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Method, r.Method))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Path, r.URL.Path))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Body, ""))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.AccessToken, accessToken))

		next.ServeHTTP(w, rCopy)
	})
}

func (m *mw) VerifyRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = lg.Ctx(r.Context(), r.Header)
			err error
		)
		defer ctx.DefError(&err)

		rCopy := r.Clone(r.Context())

		refreshToken, code := validateRefreshToken(rCopy, ctx)
		if code != codes.NoError {
			if err = lb.SendRTg(fmt.Sprintf("Ошибка в middleware при чтении refresh token. Ошибка на стороне клиента"), lb.Ranx); err != nil {
				lg.Errorf("Error send message to bot %v", err)
			}

			lg.Errorf("Error validate refresh token from RefreshToken")
			response.Write(w, codes.ErrorAccessToken, nil)
			return
		}

		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Method, r.Method))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Path, r.URL.Path))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.Body, ""))
		rCopy = rCopy.WithContext(context.WithValue(rCopy.Context(), shared.RefreshToken, refreshToken))

		next.ServeHTTP(w, rCopy)
	})
}
*/
