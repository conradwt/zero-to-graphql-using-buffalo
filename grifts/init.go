package grifts

import (
	"github.com/conradwt/zero-to-graphql-using-buffalo/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
