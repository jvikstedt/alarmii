package client

import (
	"io"

	"github.com/jvikstedt/alarmii/domain"
)

type editor interface {
	RunEditor(filePath string, initialValue []byte) (endValue []byte, err error)
}

type Client struct {
	Logger        io.Writer
	JobRepository domain.JobRepository
	Editor        editor
}
