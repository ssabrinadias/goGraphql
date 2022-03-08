package graph

import "github.com/ssabrinadias/goGraphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	list []*model.List
	user []*model.User
}
