package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/fredy-bambang/learngoyave/database/sqlboiler/model"
	"github.com/fredy-bambang/learngoyave/http/route"
	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/database"
	_ "goyave.dev/goyave/v4/database/dialect/sqlite"
)

type UserTestSuite struct { // Create a test suite for the Hello controller
	goyave.TestSuite
}

func (suite *UserTestSuite) SetupTest() {
	fmt.Println("pastiin database sudah di clear dulu")
	if err := resetTable(); err != nil {
		suite.Error(err)
	}
}

func (suite *UserTestSuite) TestRegister() {
	suite.RunServer(route.Register, func() {
		headers := map[string]string{"Content-Type": "application/json"}
		request := map[string]interface{}{
			"role":     "jack",
			"email":    "jack@example.org",
			"password": "super_Secret_password_2",
		}
		body, _ := json.Marshal(request)
		resp, err := suite.Post("/user", headers, bytes.NewReader(body))
		suite.Nil(err)
		suite.NotNil(resp)

		if resp != nil {
			defer resp.Body.Close()
			suite.Equal(http.StatusCreated, resp.StatusCode)
			json := map[string]interface{}{}
			err := suite.GetJSONBody(resp, &json)
			suite.Nil(err)
			if err == nil {
				suite.Contains(json, "id")
			}

			db, _ := database.Conn().DB()
			count, err := model.Users(model.UserWhere.Email.EQ("jack@example.org")).Count(context.Background(), db)
			if err != nil {
				suite.Error(err)
			}

			suite.Equal(int64(1), count)
		}
	})
}

func TestUserSuite(t *testing.T) {
	goyave.RunTest(t, new(UserTestSuite))
}
