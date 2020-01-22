package code

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

func ResponseErrorCode(ctx context.Context, r *http.Response, body []byte) error {
	var err error
	if r.StatusCode != http.StatusOK {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: string(body),
			Extensions: map[string]interface{}{
				"status": r.StatusCode,
			},
		})
		err = errors.New("Error of status code")
	}
	return err
 }
