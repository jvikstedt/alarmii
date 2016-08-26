package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jvikstedt/alarmii/domain"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/tylerb/graceful"
)

type Server struct {
	Port          int
	JobRepository domain.JobRepository
}

func NewServer(port int, jobRepository domain.JobRepository) Server {
	return Server{Port: port, JobRepository: jobRepository}
}

func (s Server) Start() error {
	e := echo.New()
	g := e.Group("/api/v1")
	g.POST("/jobs", s.CreateJob)
	g.GET("/jobs", s.GetAllJobs)
	g.GET("/jobs/:id", s.FindOneJob)
	std := standard.New(fmt.Sprintf(":%d", s.Port))
	std.SetHandler(e)
	err := graceful.ListenAndServe(std.Server, 5*time.Second)
	return err
}

func (s Server) GetAllJobs(c echo.Context) error {
	jobs, err := s.JobRepository.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, jobs)
}

func (s Server) FindOneJob(c echo.Context) error {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	job, err := s.JobRepository.GetOne(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, job)
}

func (s Server) CreateJob(c echo.Context) error {
	job := domain.Job{}
	if err := c.Bind(&job); err != nil {
		return err
	}
	job, err := s.JobRepository.Create(job)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusCreated, job)
}
