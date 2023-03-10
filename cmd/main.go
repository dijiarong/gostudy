package main

import "gostudy/cmd/test_echo"

// @title Swagger Example API
// @version 1.0
// @description GO测试文档.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func echoTest() {
	test_echo.StartRouter()
}

func main() {
	echoTest()
}
