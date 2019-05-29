package transport

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/api/topic"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// HTTP represents topic http service
type HTTP struct {
	svc topic.Service
}

// NewHTTP creates new topic http service
func NewHTTP(svc topic.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/topics")
	// swagger:operation POST /v1/topics topics topicCreate
	// ---
	// summary: Returns topic created.
	// description: Returns list of topics. Depending on the topic role requesting it, it may return all topics for SuperAdmin/Admin topics, all company/location topics for Company/Location admins, and an error for non-admin topics.
	// parameters:
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/topicCreate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/topicListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.POST("", h.create)

	// swagger:operation GET /v1/topics topics listTopics
	// ---
	// summary: Returns list of topics.
	// description: Returns list of topics. Depending on the topic role requesting it, it may return all topics for SuperAdmin/Admin topics, all company/location topics for Company/Location admins, and an error for non-admin topics.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/topicListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.list)

	// swagger:operation GET /v1/topics/{id} topics getTopic
	// ---
	// summary: Returns a single topic.
	// description: Returns a single topic by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of topic
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/topicResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/:id", h.view)

	// swagger:operation PATCH /v1/topics/{id} topics topicUpdate
	// ---
	// summary: Updates topic's contact information
	// description: Updates topic's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of topic
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/topicUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/topicResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/topics/{id} topics topicDelete
	// ---
	// summary: Deletes a topic
	// description: Deletes a topic with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of topic
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.DELETE("/:id", h.delete)
}

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {

		return err
	}
	usr, err := h.svc.Create(c, model.Topic{
		Name:        r.Name,
		Description: r.Description,
		Type:        r.Type,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) list(c echo.Context) error {
	p := new(model.PaginationReq)
	if err := c.Bind(p); err != nil {
		return err
	}

	result, err := h.svc.List(c, p.Transform())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result, p.Page})
}

func (h *HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}

	req := new(updateReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	usr, err := h.svc.Update(c, &topic.Update{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Mobile:    req.Mobile,
		Phone:     req.Phone,
		Address:   req.Address,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
