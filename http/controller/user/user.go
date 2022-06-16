package user

import (
	"net/http"

	"github.com/fredy-bambang/learngoyave/database/sqlboiler/model"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/database"
	"goyave.dev/goyave/v4/util/walk"
	"goyave.dev/goyave/v4/validation"
)

func Store(res *goyave.Response, req *goyave.Request) {
	db, _ := database.Conn().DB()

	isExists, err := model.Users(
		model.UserWhere.Email.EQ(req.String("email")),
	).Exists(req.Request().Context(), db)

	if err != nil {
		res.Error(err)
	}

	if isExists {
		errors := make(validation.Errors)
		errors.Add(&walk.Path{Name: "email"}, "bingung")
		errors.Add(&walk.Path{Name: "email"}, "luar biasa")

		res.JSON(
			http.StatusUnprocessableEntity,
			map[string]validation.Errors{
				"validationError": errors,
			})
		return
	}

	user := &model.User{
		ID:       uuid.New().String(),
		Email:    req.String("email"),
		Password: req.String("password"),
		Role:     req.String("role"),
	}

	if err := user.Insert(req.Request().Context(), db, boil.Infer()); err != nil {
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, user)
	}

}
