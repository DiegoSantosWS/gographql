package code

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

func erro(ctx context.Context) error {
	teste := "teste"
	gqlerror.Errorf("msg error [%s]", teste)

	graphql.AddError(ctx, &gqlerror.Error{
		Message: "alguma coisa",
		Extensions: map[string]interface{}{
			"status": "500",
		},
	})
	return nil
}
