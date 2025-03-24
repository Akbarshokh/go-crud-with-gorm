package rest

import "github.com/movie-app-crud-gorm/internal/pkg/status"

type R struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note"`
	Data      interface{} `json:"data"`
}

func view(data interface{}) R {
	return R{
		Status: status.Success,
		Data:   data,
	}
}

func errView(code int, note string) R {
	return R{
		Status:    status.Failure,
		ErrorCode: code,
		ErrorNote: note,
		Data:      nil,
	}
}
