package musiclibary

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab.com/nevasik7/lg"
)

// MusicLibraryClient - клиент для взаимодействия с внешним API.
type MusicLibraryClient struct {
	baseURL    string
	httpClient *resty.Client
}

// SongDetailResponse - структура для ответа от внешнего API.
type SongDetailResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// NewMusicLibraryClient создает новый клиент для работы с API.
func NewMusicLibraryClient(baseURL string) *MusicLibraryClient {
	return &MusicLibraryClient{
		baseURL: baseURL,
		httpClient: resty.New().
			SetHeader("Content-Type", "application/json").
			SetBaseURL(baseURL).
			EnableTrace().
			SetDebug(true),
	}
}

// GetSongDetail - получает информацию о песне из внешнего API по названию группы и песни.
// GetSongDetail запрашивает информацию о песне через внешний API.
func (c *MusicLibraryClient) GetSongDetail(group, song string) (*SongDetailResponse, error) {
	var songDetail SongDetailResponse

	lg.Infof("Запрос к API для получения данных о песне: %s - %s", group, song)

	// Выполнение GET-запроса.
	response, err := c.httpClient.R().
		SetQueryParams(map[string]string{
			"group": group,
			"song":  song,
		}).
		SetResult(&songDetail).
		Get("/info")

	if err != nil {
		return nil, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}

	// Проверка на ошибки в коде ответа.
	if response.IsError() {
		return nil, fmt.Errorf("получен некорректный статус код: %d", response.StatusCode())
	}

	lg.Infof("Получен ответ от API: %v", songDetail)
	return &songDetail, nil
}
