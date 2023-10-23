package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"

	"example/e-learn/domain"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type RoleHandler struct {
	RUsecase domain.RoleUsecase
}

// NewRoleHandler will initialize the role/ resources endpoint
func NewRoleHandler(e *echo.Echo, us domain.RoleUsecase) {
	handler := &RoleHandler{
		RUsecase: us,
	}
	// e.GET("/articles", handler.FetchArticle)
	e.POST("/roles", handler.Store)
	// e.GET("/articles/:id", handler.GetByID)
	// e.DELETE("/articles/:id", handler.Delete)
}

// FetchArticle will fetch the article based on given params
// func (a *ArticleHandler) FetchArticle(c echo.Context) error {
// 	numS := c.QueryParam("num")
// 	num, _ := strconv.Atoi(numS)
// 	cursor := c.QueryParam("cursor")
// 	ctx := c.Request().Context()

// 	listAr, nextCursor, err := a.AUsecase.Fetch(ctx, cursor, int64(num))
// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}

// 	c.Response().Header().Set(`X-Cursor`, nextCursor)
// 	return c.JSON(http.StatusOK, listAr)
// }

// GetByID will get article by given id
// func (a *ArticleHandler) GetByID(c echo.Context) error {
// 	idP, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
// 	}

// 	id := int64(idP)
// 	ctx := c.Request().Context()

// 	art, err := a.AUsecase.GetByID(ctx, id)
// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, art)
// }

func isRequestValid(m *domain.Role) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body
func (a *RoleHandler) Store(c echo.Context) (err error) {
	var role domain.Role
	err = c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&role); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.RUsecase.Store(ctx, &role)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, role)
}

// Delete will delete article by given param
// func (a *ArticleHandler) Delete(c echo.Context) error {
// 	idP, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
// 	}

// 	id := int64(idP)
// 	ctx := c.Request().Context()

// 	err = a.AUsecase.Delete(ctx, id)
// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
