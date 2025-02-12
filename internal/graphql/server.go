package graphql

// TODO move it to main API server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"

	executorsclientv1 "github.com/kubeshop/testkube-operator/client/executors/v1"
	"github.com/kubeshop/testkube/internal/graphql/graph"
	"github.com/kubeshop/testkube/pkg/event/bus"
	"github.com/kubeshop/testkube/pkg/log"
)

// const defaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	// configure NATS event bus
// 	nc, err := bus.NewNATSConnection("nats://localhost:4222")
// 	if err != nil {
// 		panic(err)
// 	}
// 	eventBus := bus.NewNATSBus(nc)

// 	kubeClient, err := kubeclient.GetClient()
// 	if err != nil {
// 		log.DefaultLogger.Panic(err)
// 	}
// 	executorsClient := executorsclientv1.NewClient(kubeClient, "testkube")

// 	srv := GetGraphQLServer(eventBus, executorsClient)

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.DefaultLogger.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.DefaultLogger.Panic(http.ListenAndServe(":"+port, nil))
// }

func GetServer(eventBus bus.Bus, executorsClient *executorsclientv1.ExecutorsClient) *handler.Server {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			Log:    log.DefaultLogger,
			Bus:    eventBus,
			Client: executorsClient,
		}}))

	srv.AddTransport(&transport.Websocket{})

	return srv

}
