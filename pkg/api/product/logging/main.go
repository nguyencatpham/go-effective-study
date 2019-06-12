package product

import (
	"github.com/nguyencatpham/go-effective-study/pkg/api/product"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// New creates new product logging service
func New(svc product.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents product logging service
type LogService struct {
	product.Service
	logger model.Logger
}

// const name = "product"

// // View logging
// func (ls *LogService) View(c echo.Context, req []string) (resp *model.Product, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "View product request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"resp": resp,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.View(c, req)
// }

// // List logging
// func (ls *LogService) List(c echo.Context, req *model.Pagination, query []string) (resp []model.Product, totalItems int, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "List product request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"resp": resp,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.List(c, req, query)
// }

// // Delete logging
// func (ls *LogService) Delete(c echo.Context, req int) (err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Delete product request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Delete(c, req)
// }

// // Update logging
// func (ls *LogService) Update(c echo.Context, req model.UpdateReq) (resp *model.Product, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Update product request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"resp": resp,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Update(c, req)
// }

// // Create logging
// func (ls *LogService) Create(c echo.Context, req model.Product) (resp *model.Product, err error) {
// 	defer func(begin time.Time) {
// 		ls.logger.Log(
// 			c,
// 			name, "Create product request", err,
// 			map[string]interface{}{
// 				"req":  req,
// 				"resp": resp,
// 				"took": time.Since(begin),
// 			},
// 		)
// 	}(time.Now())
// 	return ls.Service.Create(c, req)
// }
