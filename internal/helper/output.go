package helper

import (
	"encoding/json"
	"os"
)

type Output struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data,omitempty"`
}

func Success(data any) {
	_ = json.NewEncoder(os.Stdout).Encode(&Output{
		Success: true,
		Msg:     "ok",
		Data:    data,
	})
}

func Error(msg string) {
	_ = json.NewEncoder(os.Stdout).Encode(&Output{
		Success: false,
		Msg:     msg,
	})
}
