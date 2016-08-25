package repository_test

import (
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/jvikstedt/alarmii/domain"
	"github.com/jvikstedt/alarmii/server/repository"
	"github.com/stretchr/testify/assert"
)

var db *bolt.DB
var jobRepository *repository.JobRepository

func Setup() {
	os.Remove("test.db")
	db, _ = bolt.Open("test.db", 0600, nil)
	jobRepository = repository.NewJobRepository(db)
}

func Finish() {
	db.Close()
}

func TestMain(m *testing.M) {
	Setup()
	retCode := m.Run()
	Finish()
	os.Exit(retCode)
}

func TestCreate(t *testing.T) {
	job := domain.Job{Schedule: "@every 15s"}

	firstJob, err := jobRepository.Create(job)
	assert.Nil(t, err)
	assert.Equal(t, 1, firstJob.ID)
	assert.Equal(t, job.Schedule, firstJob.Schedule)

	job.Schedule = "@every 30s"

	secondJob, err := jobRepository.Create(job)
	assert.Nil(t, err)
	assert.Equal(t, 2, secondJob.ID)
	assert.Equal(t, job.Schedule, secondJob.Schedule)
}

func TestUpdate(t *testing.T) {
	job := domain.Job{ID: 1, Schedule: "@every 20s"}
	newJob, err := jobRepository.Update(job)
	assert.Nil(t, err)
	assert.Equal(t, 1, newJob.ID)
	assert.Equal(t, "@every 20s", newJob.Schedule)
}

func TestGetOne(t *testing.T) {
	job, err := jobRepository.GetOne(2)
	assert.Nil(t, err)
	assert.Equal(t, 2, job.ID)
	assert.Equal(t, "@every 30s", job.Schedule)

	job, err = jobRepository.GetOne(999)
	assert.Error(t, err)
	assert.Equal(t, 0, job.ID)
}

func TestDelete(t *testing.T) {
	job, err := jobRepository.Delete(2)
	assert.Nil(t, err)
	assert.Equal(t, 2, job.ID)
	assert.Equal(t, "@every 30s", job.Schedule)

	job, err = jobRepository.Delete(999)
	assert.Error(t, err)
	assert.Equal(t, 0, job.ID)
}

func TestGetAll(t *testing.T) {
	jobs, err := jobRepository.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(jobs))
	assert.Equal(t, 1, jobs[0].ID)
	assert.Equal(t, "@every 20s", jobs[0].Schedule)
}
