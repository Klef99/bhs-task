package integration

import (
	"fmt"
	"net/url"

	"github.com/Klef99/bhs-task/config"
	"github.com/Klef99/bhs-task/pkg/jwtgenerator"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/ozontech/cute"
)

type SuiteStruct struct {
	suite.Suite
	host      *url.URL
	testMaker *cute.HTTPTestMaker
	jwt       string
	cfg       *config.Config
}

func (i *SuiteStruct) BeforeAll(t provider.T) {
	// Prepare http test builder
	i.testMaker = cute.NewHTTPTestMaker()
	cfg, err := config.NewConfig("../config/config.yml")
	if err != nil {
		t.Fatalf("Config error: %v", err)
	}
	i.cfg = cfg
	// Preparing host
	host, err := url.Parse(fmt.Sprintf("http://localhost:%s/v1/", cfg.HTTP.Port))
	if err != nil {
		t.Fatalf("could not parse url, error %v", err)
	}
	i.host = host
	jtg, err := jwtgenerator.New(cfg.Jwt.Secret)
	if err != nil {
		t.Fatalf("jtg didn't create: %v", err)
	}
	jwt, err := jtg.GenerateToken("test", 1) // Basic user from test migration
	if err != nil {
		t.Fatalf("jwt didn't generate: %v", err)
	}
	i.jwt = jwt
}

func (i *SuiteStruct) BeforeEach(t provider.T) {
	t.Feature("MainTests")
}
