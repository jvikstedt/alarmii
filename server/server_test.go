package server_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/jvikstedt/alarmii/server"
	"github.com/jvikstedt/alarmii/server/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

var jobJSON = `{"schedule":"@every 15s"}`

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

func TestCreateJob(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/api/v1/jobs", strings.NewReader(jobJSON))

	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		s := server.NewServer(0, jobRepository)

		if assert.NoError(t, s.CreateJob(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, `{"id":1,"schedule":"@every 15s"}`, rec.Body.String())
		}
	}
}

func TestGetOneJob(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	s := server.NewServer(0, jobRepository)

	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/api/v1/jobs/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	if assert.NoError(t, s.FindOneJob(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"id":1,"schedule":"@every 15s"}`, rec.Body.String())
	}

	rec = httptest.NewRecorder()
	c = e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/api/v1/jobs/:id")
	c.SetParamNames("id")
	c.SetParamValues("5")
	if assert.NoError(t, s.FindOneJob(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "null", rec.Body.String())
	}
}

func TestGetAllJobs(t *testing.T) {
	e := echo.New()
	req := new(http.Request)
	s := server.NewServer(0, jobRepository)

	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/api/v1/jobs")
	if assert.NoError(t, s.GetAllJobs(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `[{"id":1,"schedule":"@every 15s"}]`, rec.Body.String())
	}
}
