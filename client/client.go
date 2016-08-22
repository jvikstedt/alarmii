package client

import (
	"io"

	"github.com/jvikstedt/alarmii/domain"
)

type Client struct {
	Logger        io.Writer
	JobRepository domain.JobRepository
}
