package integration_test

import (
	"os"
	"testing"

	"github.com/Klef99/bhs-task/integration-test"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func TestMainSuite(t *testing.T) {
	os.Setenv("ALLURE_OUTPUT_PATH", "./") // custom, read Readme.md for more info
	suite.RunSuite(t, new(integration.SuiteStruct))
}
