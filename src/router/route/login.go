package route

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = []Route{
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Function:     controllers.Login,
		AuthRequired: false,
	},
	{
		URI: "/forgot-pass",
		Method: http.MethodPost,
		Function: controllers.ForgotPass,
		AuthRequired: false,
	},
}
