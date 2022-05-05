package handler

import (
	"go-gqlgen/graph"
	"go-gqlgen/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GraphqlHandler struct{}

func NewGraphqlHandler() *GraphqlHandler {
	return &GraphqlHandler{}
}

func (h *GraphqlHandler) Playground() http.HandlerFunc {
	return playground.Handler("GraphQL playground", "/query")
}

func (h *GraphqlHandler) Query() *handler.Server {
	es := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})

	// Default pre-configured Gqlgen Handler
	// srv := handler.NewDefaultServer(es)

	// Default Gqlgen Handler without any configuration
	srv := handler.New(es)

	// Graphql Handler configuration
	// ...
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	// enable introspection if enviroment is development
	if os.Getenv("ENV") == "dev" {
		srv.Use(extension.Introspection{})
	}

	return srv
}
