package client

import (
	"encoding/json"

	"github.com/jvikstedt/alarmii/domain"
)

var jobFilePath = "job.json"

func (c Client) ListJobs() error {
	jobs, err := c.JobRepository.GetAll()
	if err != nil {
		return err
	}
	asJSON, _ := json.Marshal(jobs)
	c.Logger.Write(asJSON)
	return nil
}

func (c Client) CreateJob() error {
	job := domain.Job{}
	asPrettyJSON, _ := json.MarshalIndent(job, "", " ")
	endValue, err := c.Editor.RunEditor(jobFilePath, asPrettyJSON)
	if err != nil {
		return err
	}
	err = json.Unmarshal(endValue, &job)
	if err != nil {
		return err
	}
	job, err = c.JobRepository.Create(job)
	if err != nil {
		return err
	}
	asJSON, _ := json.Marshal(job)
	c.Logger.Write(asJSON)
	return nil
}
