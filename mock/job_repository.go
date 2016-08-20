package mock

import "github.com/jvikstedt/alarmii/domain"

type JobRepositoryMock struct {
	Receives struct {
		ID  int
		Job domain.Job
	}
	Returns struct {
		Job   domain.Job
		Jobs  []domain.Job
		Error error
	}
}

func (j *JobRepositoryMock) GetOne(id int) (domain.Job, error) {
	j.Receives.ID = id
	return j.Returns.Job, j.Returns.Error
}

func (j *JobRepositoryMock) GetAll() ([]domain.Job, error) {
	return j.Returns.Jobs, j.Returns.Error
}

func (j *JobRepositoryMock) Create(job domain.Job) (domain.Job, error) {
	j.Receives.Job = job
	return j.Returns.Job, j.Returns.Error
}

func (j *JobRepositoryMock) Update(job domain.Job) (domain.Job, error) {
	j.Receives.Job = job
	return j.Returns.Job, j.Returns.Error
}

func (j *JobRepositoryMock) Delete(id int) (domain.Job, error) {
	j.Receives.ID = id
	return j.Returns.Job, j.Returns.Error
}
