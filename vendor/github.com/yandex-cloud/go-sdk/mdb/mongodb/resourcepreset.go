// Code generated by sdkgen. DO NOT EDIT.

// nolint
package mongodb

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"
)

// ResourcePresetServiceClient is a mongodb.ResourcePresetServiceClient with
// lazy GRPC connection initialization.
type ResourcePresetServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ mongodb.ResourcePresetServiceClient = &ResourcePresetServiceClient{}

// Get implements mongodb.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) Get(ctx context.Context, in *mongodb.GetResourcePresetRequest, opts ...grpc.CallOption) (*mongodb.ResourcePreset, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewResourcePresetServiceClient(conn).Get(ctx, in, opts...)
}

// List implements mongodb.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) List(ctx context.Context, in *mongodb.ListResourcePresetsRequest, opts ...grpc.CallOption) (*mongodb.ListResourcePresetsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mongodb.NewResourcePresetServiceClient(conn).List(ctx, in, opts...)
}