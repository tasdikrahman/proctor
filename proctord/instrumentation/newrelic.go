package instrumentation

import (
	"net/http"

	"github.com/gojektech/proctor/proctord/config"
	"github.com/gojektech/proctor/proctord/logger"
	"github.com/newrelic/go-agent"
)

var NewRelicApp newrelic.Application

func InitNewRelic() error {
	if !config.NewRelicEnabled() {
		logger.Warn("New Relic is disabled")
		NewRelicApp = &StubNewrelicApp{}
		return nil
	}
	appName := config.NewRelicAppName()
	licenceKey := config.NewRelicLicenceKey()
	newRelicConfig := newrelic.NewConfig(appName, licenceKey)
	newRelicConfig.Enabled = true
	app, err := newrelic.NewApplication(newRelicConfig)
	if err != nil {
		return err
	}
	NewRelicApp = app
	return nil
}

func Wrap(pattern string, handlerFunc http.HandlerFunc) (string, func(http.ResponseWriter, *http.Request)) {
	return newrelic.WrapHandleFunc(NewRelicApp, pattern, handlerFunc)
}
