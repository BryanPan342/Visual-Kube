// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Amazon Elastic Compute Cloud (Amazon EC2) provides resizable computing capacity
// in the Amazon Web Services (AWS) cloud. Using Amazon EC2 eliminates your
// need to invest in hardware up front, so you can develop and deploy applications
// faster.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type EC2 struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "ec2"

// New creates a new instance of the EC2 client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a EC2 client from just a session.
//     svc := ec2.New(mySession)
//
//     // Create a EC2 client with additional configuration
//     svc := ec2.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *EC2 {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *EC2 {
	svc := &EC2{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-10-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(ec2query.Build)
	svc.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(ec2query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(ec2query.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a EC2 operation and runs any
// custom request initialization.
func (c *EC2) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
