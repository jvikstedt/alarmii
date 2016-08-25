package repository

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/jvikstedt/alarmii/domain"
	"github.com/jvikstedt/alarmii/util"
)

type JobRepository struct {
	db *bolt.DB
}

var jobsBucket = []byte("jobs")

func NewJobRepository(db *bolt.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (j *JobRepository) GetOne(id int) (domain.Job, error) {
	job := domain.Job{}
	err := j.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(jobsBucket)
		if b == nil {
			return errors.New("Not found")
		}
		bytes := b.Get(util.Itob(id))
		if len(bytes) == 0 {
			return errors.New("Not found")
		}
		err := json.Unmarshal(bytes, &job)
		return err
	})
	return job, err
}

func (j *JobRepository) GetAll() ([]domain.Job, error) {
	jobs := []domain.Job{}
	err := j.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(jobsBucket)
		if b == nil {
			return nil
		}
		err := b.ForEach(func(k, v []byte) error {
			var job domain.Job
			err := json.Unmarshal(v, &job)
			if err != nil {
				return err
			}
			jobs = append(jobs, job)
			return nil
		})
		return err
	})
	return jobs, err
}

func (j *JobRepository) Create(job domain.Job) (domain.Job, error) {
	err := j.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(jobsBucket)
		if err != nil {
			return err
		}
		if job.ID == 0 {
			id, err := b.NextSequence()
			if err != nil {
				return err
			}
			job.ID = int(id)
		}
		encoded, err := json.Marshal(job)
		if err != nil {
			return err
		}
		return b.Put(util.Itob(job.ID), encoded)
	})
	return job, err
}

func (j *JobRepository) Update(job domain.Job) (domain.Job, error) {
	return j.Create(job)
}

func (j *JobRepository) Delete(id int) (domain.Job, error) {
	job, err := j.GetOne(id)
	if err != nil {
		return job, err
	}
	err = j.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(jobsBucket)
		if err != nil {
			return err
		}
		err = b.Delete(util.Itob(id))
		return err
	})
	return job, err
}
