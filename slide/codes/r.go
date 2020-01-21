package codes

//Exemplo de um queryResolver
func (r *queryResolver) GetPlanets(ctx context.Context) (*planets.Planets, error) {
	req, err := http.Get("https://swapi.co/api/planets/1/")
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{Message: "alguma coisa", Extensions: map[string]interface{}{"status": "500"}})
		return nil, nil
	}
	defer req.Body.Close()
	respJSON, err := ioutil.ReadAll(req.Body)
	if err != nil {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: "alguma coisa", Extensions: map[string]interface{}{"status": "500"},
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
