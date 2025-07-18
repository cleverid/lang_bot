package errors

import (
	errors_lib "errors"
	"os"

	loggergo "github.com/nextmillenniummedia/logger-go"
)

var (
	ErrorEntityNotFound = errors_lib.New("entity not found error")
	ErrorBadRequest     = errors_lib.New("bad request")
)

func WriteErrorAndExit(err error, logger loggergo.ILogger) {
	if err == nil {
		return
	}
	logger.Error(err.Error())
	os.Exit(1)
}

func WriteErrorsAndExit(errs []error, logger loggergo.ILogger) {
	if len(errs) == 0 {
		return
	}
	for _, err := range errs {
		logger.Error(err.Error())
	}
	os.Exit(1)
}

func AppendErr(errs []error, err error) []error {
	if err == nil {
		return errs
	}
	return append(errs, err)
}
