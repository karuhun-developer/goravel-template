package tests

import (
	"github.com/goravel/framework/testing"

	"karuhundeveloper.com/gostarterkit/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
