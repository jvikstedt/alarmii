package domain

type Job struct {
	ID       int    `json:"id"`
	Schedule string `json:"schedule"`
}

type JobRepository interface {
	GetOne(id int) (Job, error)
	GetAll() ([]Job, error)
	Create(job Job) (Job, error)
	Update(job Job) (Job, error)
	Delete(id int) (Job, error)
}
