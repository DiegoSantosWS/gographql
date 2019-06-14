package manager

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/DiegoSantosWS/gographql/planets"
	"github.com/vektah/gqlparser/gqlerror"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetPlanets(ctx context.Context) (*planets.Planets, error) {
	req, err := http.Get("https://swapi.co/api/planets/1/")
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "alguma coisa",
			Extensions: map[string]interface{}{
				"status": "500",
			},
		})
		return nil, nil
	}
	defer req.Body.Close()

	respJSON, err := ioutil.ReadAll(req.Body)
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "alguma coisa",
			Extensions: map[string]interface{}{
				"status": "500",
			},
		})
		return nil, nil
	}

	pl := &planets.Planets{}
	err = json.Unmarshal(respJSON, pl)
	if err != nil {
		return nil, nil
	}

	return pl, nil
}
