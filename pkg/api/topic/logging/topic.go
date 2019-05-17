package topic

import (
	"time"

	"../../../../model"
	"github.com/labstack/echo"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/api/topic"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// New creates new topic logging service
func New(svc topic.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents topic logging service
type LogService struct {
	topic.Service
	logger model.Logger
}

const name = "topic"

// Create logging
func (ls *LogService) Create(c echo.Context, req model.Topic) (resp *model.Topic, err error) {
	defer func(begin time.Time) {
		req.Password = "xxx-redacted-xxx"
		ls.logger.Log(
			c,
			name, "Create topic request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// List logging
func (ls *LogService) List(c echo.Context, req *model.Pagination) (resp []model.Topic, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List topic request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, req)
}

// View logging
func (ls *LogService) View(c echo.Context, req int) (resp *model.Topic, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View topic request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, req)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, req int) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete topic request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, req)
}

// Update logging
func (ls *LogService) Update(c echo.Context, req *topic.Update) (resp *model.Topic, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update topic request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}
