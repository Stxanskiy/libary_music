package request

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strings"
)

// GET выполняет GET-запрос к внешнему сервису и возвращает результат через generic.
func GET[T any](client *resty.Client, params map[string]string, subURL string, headers map[string]string) (*T, error) {
	var (
		response *resty.Response
		err      error
	)

	// Если параметры заданы, преобразуем их в URL-кодированные значения.
	if len(params) > 0 {
		query := make([]string, 0, len(params))
		for key, value := range params {
			query = append(query, url.QueryEscape(key)+"="+url.QueryEscape(value))
		}
		subURL += "?" + strings.Join(query, "&")
	}

	// Выполнение запроса.
	response, err = client.R().
		SetHeaders(headers).
		Get(subURL)

	if err != nil {
		return nil, errors.New("ошибка выполнения запроса: " + err.Error())
	}

	// Проверка на успешный статус.
	if response.IsError() {
		return nil, errors.New("получен некорректный статус код: " + response.Status())
	}

	// Декодируем тело ответа в целевую структуру.
	var result T
	if err = json.Unmarshal(response.Body(), &result); err != nil {
		return nil, errors.New("ошибка декодирования ответа: " + err.Error())
	}

	return &result, nil
}
