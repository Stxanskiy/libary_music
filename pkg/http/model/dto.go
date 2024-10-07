package httpModel

type SongResponse struct {
	ID        int    `json:"id"`
	GroupName string `json:"groupName"`
	Title     string `json:"title"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
