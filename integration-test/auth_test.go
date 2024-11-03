package integration

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
)

func (i *SuiteStruct) TestLogin(t provider.T) {
	var (
		testBuilder = i.testMaker.NewTestBuilder()
	)

	u, _ := url.Parse(i.host.String())
	u.Path = path.Join(u.Path, "/login")
	testBuilder.
		Title("Login").
		Tags("one_step", "success", "json").
		Create().
		RequestBuilder(
			cute.WithURL(u),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(entity.Credentials{
				Username: "test",
				Password: "test",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("status", "Success"),
		).
		NextTest().
		Create().
		RequestBuilder(
			cute.WithURL(u),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(entity.Credentials{
				Username: "test",
				Password: "test12",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusInternalServerError).
		AssertBody(
			json.Equal("status", "error login user"),
		).
		NextTest().
		Create().
		RequestBuilder(
			cute.WithURL(u),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(entity.Credentials{
				Username: "",
				Password: "",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusInternalServerError).
		AssertBody(
			json.Equal("status", "error login user"),
		).
		ExecuteTest(context.Background(), t)
}
