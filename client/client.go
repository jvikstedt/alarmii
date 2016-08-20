package client

import (
	"encoding/json"
	"io"

	"github.com/jvikstedt/alarmii/domain"
)

type Client struct {
	Logger        io.Writer
	JobRepository domain.JobRepository
}

func (c Client) ListJobs() error {
	jobs, err := c.JobRepository.GetAll()
	if err != nil {
		return err
	}
	asJSON, _ := json.Marshal(jobs)
	c.Logger.Write(asJSON)
	return nil
}
