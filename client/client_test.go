package client_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/jvikstedt/alarmii/client"
	"github.com/jvikstedt/alarmii/domain"
	"github.com/jvikstedt/alarmii/mock"
	"github.com/stretchr/testify/assert"
)

func TestListJobs(t *testing.T) {
	var b bytes.Buffer
	jobRepositoryMock := mock.JobRepositoryMock{}
	client := client.Client{Logger: &b, JobRepository: &jobRepositoryMock}

	//Success
	jobRepositoryMock.Returns.Jobs = []domain.Job{domain.Job{ID: 1, Schedule: "@every 15s"}}
	err := client.ListJobs()
	assert.Nil(t, err)
	assert.Equal(t, `[{"id":1,"schedule":"@every 15s"}]`, b.String())

	// Error
	jobRepositoryMock.Returns.Error = errors.New("Something")
	err = client.ListJobs()
	assert.Error(t, err, "Something")
}

func TestCreateJob(t *testing.T) {
	var b bytes.Buffer
	jobRepositoryMock := mock.JobRepositoryMock{}
	client := client.Client{Logger: &b, JobRepository: &jobRepositoryMock}

	testJob := domain.Job{Schedule: "@every 10s"}

	// Success
	jobRepositoryMock.Returns.Job = domain.Job{ID: 1, Schedule: "@every 15s"}
	err := client.CreateJob(testJob)
	assert.Nil(t, err)
	assert.Equal(t, testJob, jobRepositoryMock.Receives.Job)
	assert.Equal(t, `{"id":1,"schedule":"@every 15s"}`, b.String())

	// Error
	jobRepositoryMock.Returns.Error = errors.New("Something")
	err = client.CreateJob(testJob)
	assert.Error(t, err, "Something")
}
