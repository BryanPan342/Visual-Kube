package ec2query_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil"
	"github.com/aws/aws-sdk-go/private/signer/v4"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/stretchr/testify/assert"
)

var _ bytes.Buffer // always import bytes
var _ http.Request
var _ json.Marshaler
var _ time.Time
var _ xmlutil.XMLNode
var _ xml.Attr
var _ = awstesting.GenerateAssertions
var _ = ioutil.Discard
var _ = util.Trim("")
var _ = url.Values{}
var _ = io.EOF
var _ = aws.String

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService1ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService1ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService1ProtocolTest client from just a session.
//     svc := outputservice1protocoltest.New(mySession)
//
//     // Create a OutputService1ProtocolTest client with additional configuration
//     svc := outputservice1protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService1ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService1ProtocolTest {
	c := p.ClientConfig("outputservice1protocoltest", cfgs...)
	return newOutputService1ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService1ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService1ProtocolTest {
	svc := &OutputService1ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice1protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService1ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService1ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService1TestCaseOperation1 = "OperationName"

// OutputService1TestCaseOperation1Request generates a request for the OutputService1TestCaseOperation1 operation.
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1Request(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) (req *request.Request, output *OutputService1TestShapeOutputService1TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService1TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService1TestShapeOutputService1TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService1TestShapeOutputService1TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) (*OutputService1TestShapeOutputService1TestCaseOperation1Output, error) {
	req, out := c.OutputService1TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService1TestShapeOutputService1TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService1TestShapeOutputService1TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Char *string `type:"character"`

	Double *float64 `type:"double"`

	FalseBool *bool `type:"boolean"`

	Float *float64 `type:"float"`

	Long *int64 `type:"long"`

	Num *int64 `locationName:"FooNum" type:"integer"`

	Str *string `type:"string"`

	TrueBool *bool `type:"boolean"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService2ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService2ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService2ProtocolTest client from just a session.
//     svc := outputservice2protocoltest.New(mySession)
//
//     // Create a OutputService2ProtocolTest client with additional configuration
//     svc := outputservice2protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService2ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService2ProtocolTest {
	c := p.ClientConfig("outputservice2protocoltest", cfgs...)
	return newOutputService2ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService2ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService2ProtocolTest {
	svc := &OutputService2ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice2protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService2ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService2ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService2TestCaseOperation1 = "OperationName"

// OutputService2TestCaseOperation1Request generates a request for the OutputService2TestCaseOperation1 operation.
func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1Request(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) (req *request.Request, output *OutputService2TestShapeOutputService2TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService2TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService2TestShapeOutputService2TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService2TestShapeOutputService2TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) (*OutputService2TestShapeOutputService2TestCaseOperation1Output, error) {
	req, out := c.OutputService2TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService2TestShapeOutputService2TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService2TestShapeOutputService2TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Blob []byte `type:"blob"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService3ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService3ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService3ProtocolTest client from just a session.
//     svc := outputservice3protocoltest.New(mySession)
//
//     // Create a OutputService3ProtocolTest client with additional configuration
//     svc := outputservice3protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService3ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService3ProtocolTest {
	c := p.ClientConfig("outputservice3protocoltest", cfgs...)
	return newOutputService3ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService3ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService3ProtocolTest {
	svc := &OutputService3ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice3protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService3ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService3ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService3TestCaseOperation1 = "OperationName"

// OutputService3TestCaseOperation1Request generates a request for the OutputService3TestCaseOperation1 operation.
func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1Request(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) (req *request.Request, output *OutputService3TestShapeOutputService3TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService3TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService3TestShapeOutputService3TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService3TestShapeOutputService3TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) (*OutputService3TestShapeOutputService3TestCaseOperation1Output, error) {
	req, out := c.OutputService3TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService3TestShapeOutputService3TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService3TestShapeOutputService3TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	ListMember []*string `type:"list"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService4ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService4ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService4ProtocolTest client from just a session.
//     svc := outputservice4protocoltest.New(mySession)
//
//     // Create a OutputService4ProtocolTest client with additional configuration
//     svc := outputservice4protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService4ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService4ProtocolTest {
	c := p.ClientConfig("outputservice4protocoltest", cfgs...)
	return newOutputService4ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService4ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService4ProtocolTest {
	svc := &OutputService4ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice4protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService4ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService4ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService4TestCaseOperation1 = "OperationName"

// OutputService4TestCaseOperation1Request generates a request for the OutputService4TestCaseOperation1 operation.
func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1Request(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) (req *request.Request, output *OutputService4TestShapeOutputService4TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService4TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService4TestShapeOutputService4TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService4TestShapeOutputService4TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) (*OutputService4TestShapeOutputService4TestCaseOperation1Output, error) {
	req, out := c.OutputService4TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService4TestShapeOutputService4TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService4TestShapeOutputService4TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	ListMember []*string `locationNameList:"item" type:"list"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService5ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService5ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService5ProtocolTest client from just a session.
//     svc := outputservice5protocoltest.New(mySession)
//
//     // Create a OutputService5ProtocolTest client with additional configuration
//     svc := outputservice5protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService5ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService5ProtocolTest {
	c := p.ClientConfig("outputservice5protocoltest", cfgs...)
	return newOutputService5ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService5ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService5ProtocolTest {
	svc := &OutputService5ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice5protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService5ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService5ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService5TestCaseOperation1 = "OperationName"

// OutputService5TestCaseOperation1Request generates a request for the OutputService5TestCaseOperation1 operation.
func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1Request(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) (req *request.Request, output *OutputService5TestShapeOutputService5TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService5TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService5TestShapeOutputService5TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService5TestShapeOutputService5TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) (*OutputService5TestShapeOutputService5TestCaseOperation1Output, error) {
	req, out := c.OutputService5TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService5TestShapeOutputService5TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService5TestShapeOutputService5TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	ListMember []*string `type:"list" flattened:"true"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService6ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService6ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService6ProtocolTest client from just a session.
//     svc := outputservice6protocoltest.New(mySession)
//
//     // Create a OutputService6ProtocolTest client with additional configuration
//     svc := outputservice6protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService6ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService6ProtocolTest {
	c := p.ClientConfig("outputservice6protocoltest", cfgs...)
	return newOutputService6ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService6ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService6ProtocolTest {
	svc := &OutputService6ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice6protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService6ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService6ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService6TestCaseOperation1 = "OperationName"

// OutputService6TestCaseOperation1Request generates a request for the OutputService6TestCaseOperation1 operation.
func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1Request(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) (req *request.Request, output *OutputService6TestShapeOutputService6TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService6TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService6TestShapeOutputService6TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService6TestShapeOutputService6TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) (*OutputService6TestShapeOutputService6TestCaseOperation1Output, error) {
	req, out := c.OutputService6TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService6TestShapeOutputService6TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService6TestShapeOutputService6TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Map map[string]*OutputService6TestShapeStructureType `type:"map"`
}

type OutputService6TestShapeStructureType struct {
	_ struct{} `type:"structure"`

	Foo *string `locationName:"foo" type:"string"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService7ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService7ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService7ProtocolTest client from just a session.
//     svc := outputservice7protocoltest.New(mySession)
//
//     // Create a OutputService7ProtocolTest client with additional configuration
//     svc := outputservice7protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService7ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService7ProtocolTest {
	c := p.ClientConfig("outputservice7protocoltest", cfgs...)
	return newOutputService7ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService7ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService7ProtocolTest {
	svc := &OutputService7ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice7protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService7ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService7ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService7TestCaseOperation1 = "OperationName"

// OutputService7TestCaseOperation1Request generates a request for the OutputService7TestCaseOperation1 operation.
func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1Request(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) (req *request.Request, output *OutputService7TestShapeOutputService7TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService7TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService7TestShapeOutputService7TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService7TestShapeOutputService7TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) (*OutputService7TestShapeOutputService7TestCaseOperation1Output, error) {
	req, out := c.OutputService7TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService7TestShapeOutputService7TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService7TestShapeOutputService7TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Map map[string]*string `type:"map" flattened:"true"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService8ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService8ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService8ProtocolTest client from just a session.
//     svc := outputservice8protocoltest.New(mySession)
//
//     // Create a OutputService8ProtocolTest client with additional configuration
//     svc := outputservice8protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService8ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService8ProtocolTest {
	c := p.ClientConfig("outputservice8protocoltest", cfgs...)
	return newOutputService8ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService8ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService8ProtocolTest {
	svc := &OutputService8ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice8protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService8ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService8ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService8TestCaseOperation1 = "OperationName"

// OutputService8TestCaseOperation1Request generates a request for the OutputService8TestCaseOperation1 operation.
func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1Request(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) (req *request.Request, output *OutputService8TestShapeOutputService8TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService8TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService8TestShapeOutputService8TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService8TestShapeOutputService8TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) (*OutputService8TestShapeOutputService8TestCaseOperation1Output, error) {
	req, out := c.OutputService8TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService8TestShapeOutputService8TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService8TestShapeOutputService8TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Map map[string]*string `locationNameKey:"foo" locationNameValue:"bar" type:"map" flattened:"true"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type OutputService9ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the OutputService9ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService9ProtocolTest client from just a session.
//     svc := outputservice9protocoltest.New(mySession)
//
//     // Create a OutputService9ProtocolTest client with additional configuration
//     svc := outputservice9protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService9ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *OutputService9ProtocolTest {
	c := p.ClientConfig("outputservice9protocoltest", cfgs...)
	return newOutputService9ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newOutputService9ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *OutputService9ProtocolTest {
	svc := &OutputService9ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "outputservice9protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "",
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

	return svc
}

// newRequest creates a new request for a OutputService9ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService9ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService9TestCaseOperation1 = "OperationName"

// OutputService9TestCaseOperation1Request generates a request for the OutputService9TestCaseOperation1 operation.
func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1Request(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) (req *request.Request, output *OutputService9TestShapeOutputService9TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opOutputService9TestCaseOperation1,
	}

	if input == nil {
		input = &OutputService9TestShapeOutputService9TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &OutputService9TestShapeOutputService9TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) (*OutputService9TestShapeOutputService9TestCaseOperation1Output, error) {
	req, out := c.OutputService9TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type OutputService9TestShapeOutputService9TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService9TestShapeOutputService9TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Foo *string `type:"string"`
}

//
// Tests begin here
//

func TestOutputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService1ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><Str>myname</Str><FooNum>123</FooNum><FalseBool>false</FalseBool><TrueBool>true</TrueBool><Float>1.2</Float><Double>1.3</Double><Long>200</Long><Char>a</Char><RequestId>request-id</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService1TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "a", *out.Char)
	assert.Equal(t, 1.3, *out.Double)
	assert.Equal(t, false, *out.FalseBool)
	assert.Equal(t, 1.2, *out.Float)
	assert.Equal(t, int64(200), *out.Long)
	assert.Equal(t, int64(123), *out.Num)
	assert.Equal(t, "myname", *out.Str)
	assert.Equal(t, true, *out.TrueBool)

}

func TestOutputService2ProtocolTestBlobCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService2ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><Blob>dmFsdWU=</Blob><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService2TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "value", string(out.Blob))

}

func TestOutputService3ProtocolTestListsCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService3ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><ListMember><member>abc</member><member>123</member></ListMember><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService3TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.ListMember[0])
	assert.Equal(t, "123", *out.ListMember[1])

}

func TestOutputService4ProtocolTestListWithCustomMemberNameCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService4ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><ListMember><item>abc</item><item>123</item></ListMember><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService4TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.ListMember[0])
	assert.Equal(t, "123", *out.ListMember[1])

}

func TestOutputService5ProtocolTestFlattenedListCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService5ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><ListMember>abc</ListMember><ListMember>123</ListMember><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService5TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "abc", *out.ListMember[0])
	assert.Equal(t, "123", *out.ListMember[1])

}

func TestOutputService6ProtocolTestNormalMapCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService6ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><Map><entry><key>qux</key><value><foo>bar</foo></value></entry><entry><key>baz</key><value><foo>bam</foo></value></entry></Map><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService6TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "bam", *out.Map["baz"].Foo)
	assert.Equal(t, "bar", *out.Map["qux"].Foo)

}

func TestOutputService7ProtocolTestFlattenedMapCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService7ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><Map><key>qux</key><value>bar</value></Map><Map><key>baz</key><value>bam</value></Map><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService7TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "bam", *out.Map["baz"])
	assert.Equal(t, "bar", *out.Map["qux"])

}

func TestOutputService8ProtocolTestNamedMapCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService8ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><Map><foo>qux</foo><bar>bar</bar></Map><Map><foo>baz</foo><bar>bam</bar></Map><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService8TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "bam", *out.Map["baz"])
	assert.Equal(t, "bar", *out.Map["qux"])

}

func TestOutputService9ProtocolTestEmptyStringCase1(t *testing.T) {
	sess := session.New()
	svc := NewOutputService9ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	buf := bytes.NewReader([]byte("<OperationNameResponse><Foo/><RequestId>requestid</RequestId></OperationNameResponse>"))
	req, out := svc.OutputService9TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	ec2query.UnmarshalMeta(req)
	ec2query.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	assert.Equal(t, "", *out.Foo)

}
