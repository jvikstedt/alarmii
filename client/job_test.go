package client_test

import (
	"bytes"
	"encoding/json"
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
	assert.EqualError(t, err, "Something")
}

func TestCreateJob(t *testing.T) {
	var b bytes.Buffer
	jobRepositoryMock := mock.JobRepositoryMock{}
	editorMock := mock.EditorMock{}
	client := client.Client{Logger: &b, JobRepository: &jobRepositoryMock, Editor: &editorMock}

	// Editor Error
	editorMock.Returns.Error = errors.New("EditorError")
	err := client.CreateJob()
	assert.EqualError(t, err, "EditorError")

	// Prettyfied json
	job := domain.Job{}
	asPrettyJSON, _ := json.MarshalIndent(job, "", " ")
	client.CreateJob()
	assert.Equal(t, asPrettyJSON, editorMock.Receives.InitialValue)

	// Editor invalid json
	editorMock.Returns.Error = nil
	editorMock.Returns.EndValue = []byte("blablabla")
	err = client.CreateJob()
	assert.Contains(t, err.Error(), "invalid character")

	// Success
	editorMock.Returns.EndValue = []byte(`{"id":1,"schedule":"@every 15s"}`)
	jobRepositoryMock.Returns.Job = domain.Job{ID: 1, Schedule: "@every 15s"}
	err = client.CreateJob()
	assert.Equal(t, jobRepositoryMock.Receives.Job.ID, 1)
	assert.Equal(t, jobRepositoryMock.Receives.Job.Schedule, "@every 15s")
	assert.Equal(t, `{"id":1,"schedule":"@every 15s"}`, b.String())

	// Error
	jobRepositoryMock.Returns.Error = errors.New("CreateError")
	err = client.CreateJob()
	assert.EqualError(t, err, "CreateError")
}
