package codes

import (
	"net/http"
)

type InternalErrorDto struct {
	Message string
	Status  int
}

const (
	UnknownError = -1
	NoError      = 0

	// Unauthorized general error code
	Unauthorized           = 10000
	InvalidPasswordOrEmail = 10001
	InvalidSign            = 10002

	// InputError general error code
	InputError           = 20000
	BodyBinding          = 20001
	QueryBinding         = 20002
	ParsingKey           = 20003
	SignHeader           = 20004
	MaximumSessions      = 20008
	EmptyEmailOrPassword = 20009
	InvalidEmail         = 20010
	InvalidPassword      = 20011
	ErrorReadQueryParam  = 20012
	UnknownMonth         = 20013
	CuratorNotFound      = 20014
	NotValidYear         = 20015
	ErrorParsingDate     = 20016

	// Forbidden general error code
	Forbidden = 30000

	// InternalError general error code
	BodyRead                            = 40010
	PublicKeyNotFound                   = 40015
	ErrorInsertActivity                 = 40016
	ErrorGetActivityByCuratorId         = 40018
	ErrorDeleteActivity                 = 40019
	ErrorGetRoleByActivityID            = 40020
	ErrorNotRights                      = 40021
	ErrorGetActivityByData              = 40022
	ErrorUpdateActivity                 = 40023
	ErrorGetActivityByID                = 40024
	ErrorGetCountActivityGroupByCurator = 40025
	ErrorSendQueryToBitrixAPI           = 40026
	ErrorSendQueryToRanepaAPI           = 40027
	UserDisactive                       = 40028
	ErrorAccessToken                    = 40029
	NotFoundRefreshToken                = 40030
	NotFoundTitleCookies                = 40031
)

var InternalErr = map[int]InternalErrorDto{
	// 1XXXX credentials validation errors
	Unauthorized: {
		Message: "unauthorized",
		Status:  http.StatusUnauthorized,
	},
	InvalidPasswordOrEmail: {
		Message: "invalid email or password",
		Status:  http.StatusUnauthorized,
	},
	InvalidSign: {
		Message: "invalid signature or pubKey",
		Status:  http.StatusUnauthorized,
	},
	PublicKeyNotFound: {
		Message: "invalid public_key",
		Status:  http.StatusUnauthorized,
	},
	ErrorNotRights: {
		Message: "",
		Status:  http.StatusUnauthorized,
	},
	NotFoundRefreshToken: {
		Message: "not found refresh_token",
		Status:  http.StatusUnauthorized,
	},
	NotFoundTitleCookies: {
		Message: "not found title refresh_token from cookie",
		Status:  http.StatusUnauthorized,
	},

	// 2XXXX errors caused by user input
	InputError: {
		Message: "wrong input",
		Status:  http.StatusBadRequest,
	},
	BodyBinding: {
		Message: "can't bind body to request model",
		Status:  http.StatusUnprocessableEntity,
	},
	QueryBinding: {
		Message: "can't bind query parameters",
		Status:  http.StatusUnprocessableEntity,
	},
	ParsingKey: {
		Message: "failed to parse key",
		Status:  http.StatusUnprocessableEntity,
	},
	SignHeader: {
		Message: "signature header value missing or malformed",
		Status:  http.StatusBadRequest,
	},
	MaximumSessions: {
		Message: "maximum number of sessions is already created",
		Status:  http.StatusBadRequest,
	},
	EmptyEmailOrPassword: {
		Message: "email or password is empty",
		Status:  http.StatusBadRequest,
	},
	InvalidEmail: {
		Message: "invalid email",
		Status:  http.StatusBadRequest,
	},
	InvalidPassword: {
		Message: "invalid password",
		Status:  http.StatusBadRequest,
	},
	ErrorReadQueryParam: {
		Message: "error read query param",
		Status:  http.StatusBadRequest,
	},
	UnknownMonth: {
		Message: "unknown numeric month",
		Status:  http.StatusBadRequest,
	},
	CuratorNotFound: {
		Message: "curatorID not found",
		Status:  http.StatusBadRequest,
	},
	NotValidYear: {
		Message: "year not valid",
		Status:  http.StatusBadRequest,
	},

	ErrorParsingDate: {
		Message: "error, date not valid",
		Status:  http.StatusBadRequest,
	},

	// 3XXXX no access to resources
	Forbidden: {
		Message: "forbidden",
		Status:  http.StatusForbidden,
	},
	// 4XXXX internal errors
	UnknownError: {
		Message: "unknown internal error",
		Status:  http.StatusInternalServerError,
	},
	ErrorInsertActivity: {
		Message: "error insert to activity table",
		Status:  http.StatusInternalServerError,
	},
	ErrorGetActivityByCuratorId: {
		Message: "error get all by curatorID",
		Status:  http.StatusInternalServerError,
	},
	ErrorDeleteActivity: {
		Message: "error delete activity",
		Status:  http.StatusInternalServerError,
	},
	ErrorGetRoleByActivityID: {
		Message: "error get role from db by activityID",
		Status:  http.StatusInternalServerError,
	},
	ErrorGetActivityByData: {
		Message: "error get activity by month",
		Status:  http.StatusInternalServerError,
	},
	ErrorUpdateActivity: {
		Message: "error update activity",
		Status:  http.StatusInternalServerError,
	},
	ErrorGetActivityByID: {
		Message: "error get activity by ID",
		Status:  http.StatusInternalServerError,
	},
	ErrorGetCountActivityGroupByCurator: {
		Message: "error get count activity by curatorID",
		Status:  http.StatusInternalServerError,
	},
	ErrorSendQueryToBitrixAPI: {
		Message: "error get token to api bitrix",
		Status:  http.StatusInternalServerError,
	},
	ErrorSendQueryToRanepaAPI: {
		Message: "error get client to api musiclibary",
		Status:  http.StatusInternalServerError,
	},
	UserDisactive: {
		Message: "error user disactive",
		Status:  http.StatusInternalServerError,
	},
	ErrorAccessToken: {
		Message: "error read access token from header",
		Status:  http.StatusForbidden,
	},
}
