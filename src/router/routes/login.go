package routes

import (
	"net/http"
	"user-api/src/controllers"
)

var loginRoute = Route{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	AuthenticationRequired: false,
}
