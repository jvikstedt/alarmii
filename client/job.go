package client

import (
	"encoding/json"

	"github.com/jvikstedt/alarmii/domain"
)

func (c Client) ListJobs() error {
	jobs, err := c.JobRepository.GetAll()
	if err != nil {
		return err
	}
	asJSON, _ := json.Marshal(jobs)
	c.Logger.Write(asJSON)
	return nil
}

func (c Client) CreateJob(job domain.Job) error {
	job, err := c.JobRepository.Create(job)
	if err != nil {
		return err
	}
	asJSON, _ := json.Marshal(job)
	c.Logger.Write(asJSON)
	return nil
}
