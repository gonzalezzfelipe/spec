// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: utxorpc/watch/v1/watch.proto

package watchv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/bufbuild/buf-tour/gen/utxorpc/watch/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ChainWatchServiceName is the fully-qualified name of the ChainWatchService service.
	ChainWatchServiceName = "utxorpc.watch.v1.ChainWatchService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChainWatchServiceWaitForTxProcedure is the fully-qualified name of the ChainWatchService's
	// WaitForTx RPC.
	ChainWatchServiceWaitForTxProcedure = "/utxorpc.watch.v1.ChainWatchService/WaitForTx"
	// ChainWatchServiceFollowTxsProcedure is the fully-qualified name of the ChainWatchService's
	// FollowTxs RPC.
	ChainWatchServiceFollowTxsProcedure = "/utxorpc.watch.v1.ChainWatchService/FollowTxs"
)

// ChainWatchServiceClient is a client for the utxorpc.watch.v1.ChainWatchService service.
type ChainWatchServiceClient interface {
	WaitForTx(context.Context, *connect_go.Request[v1.WaitForTxRequest]) (*connect_go.Response[v1.WaitForTxResponse], error)
	FollowTxs(context.Context, *connect_go.Request[v1.FollowTxsRequest]) (*connect_go.ServerStreamForClient[v1.FollowTxsResponse], error)
}

// NewChainWatchServiceClient constructs a client for the utxorpc.watch.v1.ChainWatchService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChainWatchServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ChainWatchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chainWatchServiceClient{
		waitForTx: connect_go.NewClient[v1.WaitForTxRequest, v1.WaitForTxResponse](
			httpClient,
			baseURL+ChainWatchServiceWaitForTxProcedure,
			opts...,
		),
		followTxs: connect_go.NewClient[v1.FollowTxsRequest, v1.FollowTxsResponse](
			httpClient,
			baseURL+ChainWatchServiceFollowTxsProcedure,
			opts...,
		),
	}
}

// chainWatchServiceClient implements ChainWatchServiceClient.
type chainWatchServiceClient struct {
	waitForTx *connect_go.Client[v1.WaitForTxRequest, v1.WaitForTxResponse]
	followTxs *connect_go.Client[v1.FollowTxsRequest, v1.FollowTxsResponse]
}

// WaitForTx calls utxorpc.watch.v1.ChainWatchService.WaitForTx.
func (c *chainWatchServiceClient) WaitForTx(ctx context.Context, req *connect_go.Request[v1.WaitForTxRequest]) (*connect_go.Response[v1.WaitForTxResponse], error) {
	return c.waitForTx.CallUnary(ctx, req)
}

// FollowTxs calls utxorpc.watch.v1.ChainWatchService.FollowTxs.
func (c *chainWatchServiceClient) FollowTxs(ctx context.Context, req *connect_go.Request[v1.FollowTxsRequest]) (*connect_go.ServerStreamForClient[v1.FollowTxsResponse], error) {
	return c.followTxs.CallServerStream(ctx, req)
}

// ChainWatchServiceHandler is an implementation of the utxorpc.watch.v1.ChainWatchService service.
type ChainWatchServiceHandler interface {
	WaitForTx(context.Context, *connect_go.Request[v1.WaitForTxRequest]) (*connect_go.Response[v1.WaitForTxResponse], error)
	FollowTxs(context.Context, *connect_go.Request[v1.FollowTxsRequest], *connect_go.ServerStream[v1.FollowTxsResponse]) error
}

// NewChainWatchServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChainWatchServiceHandler(svc ChainWatchServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ChainWatchServiceWaitForTxProcedure, connect_go.NewUnaryHandler(
		ChainWatchServiceWaitForTxProcedure,
		svc.WaitForTx,
		opts...,
	))
	mux.Handle(ChainWatchServiceFollowTxsProcedure, connect_go.NewServerStreamHandler(
		ChainWatchServiceFollowTxsProcedure,
		svc.FollowTxs,
		opts...,
	))
	return "/utxorpc.watch.v1.ChainWatchService/", mux
}

// UnimplementedChainWatchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChainWatchServiceHandler struct{}

func (UnimplementedChainWatchServiceHandler) WaitForTx(context.Context, *connect_go.Request[v1.WaitForTxRequest]) (*connect_go.Response[v1.WaitForTxResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("utxorpc.watch.v1.ChainWatchService.WaitForTx is not implemented"))
}

func (UnimplementedChainWatchServiceHandler) FollowTxs(context.Context, *connect_go.Request[v1.FollowTxsRequest], *connect_go.ServerStream[v1.FollowTxsResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("utxorpc.watch.v1.ChainWatchService.FollowTxs is not implemented"))
}