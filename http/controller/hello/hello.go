package hello

import (
	"net/http"

	"github.com/fredy-bambang/learngoyave/database/sqlboiler/model"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/database"
)

// Controllers are files containing a collection of Handlers related to a specific feature.
// Each feature should have its own package.
//
// Controller handlers contain the business logic of your application.
// They should be concise and focused on what matters for this particular feature in your application.
// Learn more about controllers here: https://goyave.dev/guide/basics/controllers.html

// ----------------------------------------------------------------------

// SayHi is a controller handler writing "Hi!" as a response.
//
// The Response object is used to write your response.
// https://goyave.dev/guide/basics/responses.html
//
// The Request object contains all the information about the incoming request, including it's parsed body,
// query params and route parameters.
// https://goyave.dev/guide/basics/requests.html
func SayHi(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, "Hi!")
}

// Echo is a controller handler writing the input field "text" as a response.
// This route is validated. See "http/request/echorequest/echo.go" for more details.
func Echo(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, request.String("text"))
}

// StoreUser .
func StoreUser(response *goyave.Response, request *goyave.Request) {
	db, _ := database.Conn().DB()
	task := &model.Task{
		ID:         uuid.New().String(),
		Title:      null.StringFrom("test title"),
		IsFinished: null.BoolFrom(true),
	}

	if err := task.Insert(request.Request().Context(), db, boil.Infer()); err != nil {
		response.Error(err)
	}

	taskOutput, _ := model.Tasks().All(request.Request().Context(), db)
	response.JSON(http.StatusOK, taskOutput)
}
