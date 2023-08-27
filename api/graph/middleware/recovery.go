package middleware

import (
	"context"
	"errors"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func errorType(errAny interface{}) error {
	var err error
	switch t := errAny.(type) {
	case string:
		err = errors.New(t)
	case error:
		err = t
	default:
		err = errors.New("unknown error")
	}
	return err
}

func GraphQLRecovery(ctx context.Context, r any) error {
	err := errorType(r)
	log.Error().Msg(err.Error())
	debug.PrintStack()
	return gqlerror.Errorf("internal server error")
}
