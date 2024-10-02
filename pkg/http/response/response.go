package response

import (
	"encoding/json"
	"gitlab.com/nevasik7/lg"
	"net/http"
)

// BaseResponse определяет структуру стандартного ответа API.
type BaseResponse struct {
	Success bool        `json:"success"`         // Индикатор успешности
	Data    any         `json:"data,omitempty"`  // Данные, возвращаемые клиенту
	Error   *ErrorModel `json:"error,omitempty"` // Модель ошибки
}

// ErrorModel содержит данные об ошибке.
type ErrorModel struct {
	Code    int    `json:"code"`    // Код ошибки
	Message string `json:"message"` // Описание ошибки
}

// Write формирует и отправляет ответ в формате JSON.
func Write(w http.ResponseWriter, code int, data any, errMessage string) {
	// Устанавливаем тип контента.
	w.Header().Set("Content-Type", "application/json")

	// Формируем структуру ответа.
	var resp BaseResponse
	resp.Success = code == http.StatusOK
	resp.Data = data

	// Если код не соответствует успешному статусу, добавляем ошибку.
	if !resp.Success {
		resp.Error = &ErrorModel{
			Code:    code,
			Message: errMessage,
		}
		w.WriteHeader(code)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	// Преобразуем ответ в JSON.
	responseBytes, err := json.Marshal(resp)
	if err != nil {
		lg.Errorf("ошибка при сериализации ответа: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту.
	if _, err := w.Write(responseBytes); err != nil {
		lg.Errorf("ошибка при отправке ответа: %v", err)
	}
}

// WriteSuccess отправляет успешный ответ клиенту.
func WriteSuccess(w http.ResponseWriter, data any) {
	Write(w, http.StatusOK, data, "")
}

// WriteError отправляет ошибку с заданным кодом и сообщением.
func WriteError(w http.ResponseWriter, code int, message string) {
	Write(w, code, nil, message)
}
