package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type apiServer struct {
	productCatalogSvcAddr string
	productCatalogSvcConn *grpc.ClientConn
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}

func main() {
	ctx := context.Background()
	svc := new(apiServer)
	mustConnGRPC(ctx, &svc.productCatalogSvcConn, svc.productCatalogSvcAddr)

	router := mux.NewRouter()
	router.HandleFunc("/products", svc.productsHandler).Methods(http.MethodGet, http.MethodHead)
	//router.Handle("/products", handlers.CreateProductHandler()).Methods("POST")
	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	fmt.Println("Staring Product Catalog server on Port 9090")
	server.ListenAndServe()
}
