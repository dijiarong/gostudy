package test_echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// -----------------------------------API HelloWorldHandler
type HelloWorldHandler struct{}

// @Summary Test
// @Description 输出hello world
// @Tags Test
// @Accept json
// @Produce json
// @Success 200 {string} string "Success"
// @Failure 400 {string} string "Failed"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Error"
// @Router /hello-world [get]
func (receiver HelloWorldHandler) main(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
func (receiver HelloWorldHandler) router() {
	GetEcho().Add("GET", "/hello-world", receiver.main)
}

// -----------------------------------API PrintName
type PrintName struct{}
type PrintNameResponse struct {
	Name string `json:"name"`
}

// @Summary Test
// @Description 根据参数输出name
// @Tags Test
// @Accept json
// @Produce json
// @Param name query string true "名称"
// @Success 200 {object} PrintNameResponse PrintNameResponse{}
// @Failure 400 {string} string "Failed"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Error"
// @Router /print-name [get]
func (receiver PrintName) main(c echo.Context) error {
	name := c.QueryParam("name") // 获取query参数
	return c.JSON(http.StatusOK, PrintNameResponse{Name: name})
}
func (receiver PrintName) router() {
	GetEcho().Add("GET", "/print-name", receiver.main)
}
