package client_test

import (
	"bytes"
	"testing"

	"github.com/jvikstedt/alarmii/client"
	"github.com/jvikstedt/alarmii/domain"
	"github.com/jvikstedt/alarmii/mock"
	"github.com/stretchr/testify/assert"
)

func TestListJobs(t *testing.T) {
	var b bytes.Buffer
	jobRepositoryMock := mock.JobRepositoryMock{}
	jobRepositoryMock.Returns.Jobs = []domain.Job{domain.Job{ID: 1, Schedule: "@every 15s"}}
	client := client.Client{Logger: &b, JobRepository: &jobRepositoryMock}
	err := client.ListJobs()
	assert.Nil(t, err)
	assert.Equal(t, `[{"id":1,"schedule":"@every 15s"}]`, b.String())
}
