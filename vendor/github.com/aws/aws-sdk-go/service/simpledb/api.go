// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package simpledb provides a client for Amazon SimpleDB.
package simpledb

import (
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opBatchDeleteAttributes = "BatchDeleteAttributes"

// BatchDeleteAttributesRequest generates a request for the BatchDeleteAttributes operation.
func (c *SimpleDB) BatchDeleteAttributesRequest(input *BatchDeleteAttributesInput) (req *request.Request, output *BatchDeleteAttributesOutput) {
	op := &request.Operation{
		Name:       opBatchDeleteAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeleteAttributesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &BatchDeleteAttributesOutput{}
	req.Data = output
	return
}

// Performs multiple DeleteAttributes operations in a single call, which reduces
// round trips and latencies. This enables Amazon SimpleDB to optimize requests,
// which generally yields better throughput.
//
//   If you specify BatchDeleteAttributes without attributes or values, all
// the attributes for the item are deleted.
//
//  BatchDeleteAttributes is an idempotent operation; running it multiple times
// on the same item or attribute doesn't result in an error.
//
//  The BatchDeleteAttributes operation succeeds or fails in its entirety.
// There are no partial deletes. You can execute multiple BatchDeleteAttributes
// operations and other operations in parallel. However, large numbers of concurrent
// BatchDeleteAttributes calls can result in Service Unavailable (503) responses.
//
//  This operation is vulnerable to exceeding the maximum URL size when making
// a REST request using the HTTP GET method.
//
//  This operation does not support conditions using Expected.X.Name, Expected.X.Value,
// or Expected.X.Exists.
//
//   The following limitations are enforced for this operation:  1 MB request
// size 25 item limit per BatchDeleteAttributes operation
func (c *SimpleDB) BatchDeleteAttributes(input *BatchDeleteAttributesInput) (*BatchDeleteAttributesOutput, error) {
	req, out := c.BatchDeleteAttributesRequest(input)
	err := req.Send()
	return out, err
}

const opBatchPutAttributes = "BatchPutAttributes"

// BatchPutAttributesRequest generates a request for the BatchPutAttributes operation.
func (c *SimpleDB) BatchPutAttributesRequest(input *BatchPutAttributesInput) (req *request.Request, output *BatchPutAttributesOutput) {
	op := &request.Operation{
		Name:       opBatchPutAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchPutAttributesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &BatchPutAttributesOutput{}
	req.Data = output
	return
}

// The BatchPutAttributes operation creates or replaces attributes within one
// or more items. By using this operation, the client can perform multiple PutAttribute
// operation with a single call. This helps yield savings in round trips and
// latencies, enabling Amazon SimpleDB to optimize requests and generally produce
// better throughput.
//
//  The client may specify the item name with the Item.X.ItemName parameter.
// The client may specify new attributes using a combination of the Item.X.Attribute.Y.Name
// and Item.X.Attribute.Y.Value parameters. The client may specify the first
// attribute for the first item using the parameters Item.0.Attribute.0.Name
// and Item.0.Attribute.0.Value, and for the second attribute for the first
// item by the parameters Item.0.Attribute.1.Name and Item.0.Attribute.1.Value,
// and so on.
//
//  Attributes are uniquely identified within an item by their name/value combination.
// For example, a single item can have the attributes { "first_name", "first_value"
// } and { "first_name", "second_value" }. However, it cannot have two attribute
// instances where both the Item.X.Attribute.Y.Name and Item.X.Attribute.Y.Value
// are the same.
//
//  Optionally, the requester can supply the Replace parameter for each individual
// value. Setting this value to true will cause the new attribute values to
// replace the existing attribute values. For example, if an item I has the
// attributes { 'a', '1' }, { 'b', '2'} and { 'b', '3' } and the requester does
// a BatchPutAttributes of {'I', 'b', '4' } with the Replace parameter set to
// true, the final attributes of the item will be { 'a', '1' } and { 'b', '4'
// }, replacing the previous values of the 'b' attribute with the new value.
//
//  You cannot specify an empty string as an item or as an attribute name.
// The BatchPutAttributes operation succeeds or fails in its entirety. There
// are no partial puts.   This operation is vulnerable to exceeding the maximum
// URL size when making a REST request using the HTTP GET method. This operation
// does not support conditions using Expected.X.Name, Expected.X.Value, or Expected.X.Exists.
//   You can execute multiple BatchPutAttributes operations and other operations
// in parallel. However, large numbers of concurrent BatchPutAttributes calls
// can result in Service Unavailable (503) responses.
//
//  The following limitations are enforced for this operation:  256 attribute
// name-value pairs per item 1 MB request size 1 billion attributes per domain
// 10 GB of total user data storage per domain 25 item limit per BatchPutAttributes
// operation
func (c *SimpleDB) BatchPutAttributes(input *BatchPutAttributesInput) (*BatchPutAttributesOutput, error) {
	req, out := c.BatchPutAttributesRequest(input)
	err := req.Send()
	return out, err
}

const opCreateDomain = "CreateDomain"

// CreateDomainRequest generates a request for the CreateDomain operation.
func (c *SimpleDB) CreateDomainRequest(input *CreateDomainInput) (req *request.Request, output *CreateDomainOutput) {
	op := &request.Operation{
		Name:       opCreateDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDomainInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateDomainOutput{}
	req.Data = output
	return
}

// The CreateDomain operation creates a new domain. The domain name should be
// unique among the domains associated with the Access Key ID provided in the
// request. The CreateDomain operation may take 10 or more seconds to complete.
//
//  CreateDomain is an idempotent operation; running it multiple times using
// the same domain name will not result in an error response.   The client can
// create up to 100 domains per account.
//
//  If the client requires additional domains, go to  http://aws.amazon.com/contact-us/simpledb-limit-request/
// (http://aws.amazon.com/contact-us/simpledb-limit-request/).
func (c *SimpleDB) CreateDomain(input *CreateDomainInput) (*CreateDomainOutput, error) {
	req, out := c.CreateDomainRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteAttributes = "DeleteAttributes"

// DeleteAttributesRequest generates a request for the DeleteAttributes operation.
func (c *SimpleDB) DeleteAttributesRequest(input *DeleteAttributesInput) (req *request.Request, output *DeleteAttributesOutput) {
	op := &request.Operation{
		Name:       opDeleteAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAttributesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteAttributesOutput{}
	req.Data = output
	return
}

// Deletes one or more attributes associated with an item. If all attributes
// of the item are deleted, the item is deleted.
//
//  If DeleteAttributes is called without being passed any attributes or values
// specified, all the attributes for the item are deleted.   DeleteAttributes
// is an idempotent operation; running it multiple times on the same item or
// attribute does not result in an error response.
//
//  Because Amazon SimpleDB makes multiple copies of item data and uses an
// eventual consistency update model, performing a GetAttributes or Select operation
// (read) immediately after a DeleteAttributes or PutAttributes operation (write)
// might not return updated item data.
func (c *SimpleDB) DeleteAttributes(input *DeleteAttributesInput) (*DeleteAttributesOutput, error) {
	req, out := c.DeleteAttributesRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteDomain = "DeleteDomain"

// DeleteDomainRequest generates a request for the DeleteDomain operation.
func (c *SimpleDB) DeleteDomainRequest(input *DeleteDomainInput) (req *request.Request, output *DeleteDomainOutput) {
	op := &request.Operation{
		Name:       opDeleteDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDomainInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteDomainOutput{}
	req.Data = output
	return
}

// The DeleteDomain operation deletes a domain. Any items (and their attributes)
// in the domain are deleted as well. The DeleteDomain operation might take
// 10 or more seconds to complete.
//
//  Running DeleteDomain on a domain that does not exist or running the function
// multiple times using the same domain name will not result in an error response.
func (c *SimpleDB) DeleteDomain(input *DeleteDomainInput) (*DeleteDomainOutput, error) {
	req, out := c.DeleteDomainRequest(input)
	err := req.Send()
	return out, err
}

const opDomainMetadata = "DomainMetadata"

// DomainMetadataRequest generates a request for the DomainMetadata operation.
func (c *SimpleDB) DomainMetadataRequest(input *DomainMetadataInput) (req *request.Request, output *DomainMetadataOutput) {
	op := &request.Operation{
		Name:       opDomainMetadata,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DomainMetadataInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DomainMetadataOutput{}
	req.Data = output
	return
}

// Returns information about the domain, including when the domain was created,
// the number of items and attributes in the domain, and the size of the attribute
// names and values.
func (c *SimpleDB) DomainMetadata(input *DomainMetadataInput) (*DomainMetadataOutput, error) {
	req, out := c.DomainMetadataRequest(input)
	err := req.Send()
	return out, err
}

const opGetAttributes = "GetAttributes"

// GetAttributesRequest generates a request for the GetAttributes operation.
func (c *SimpleDB) GetAttributesRequest(input *GetAttributesInput) (req *request.Request, output *GetAttributesOutput) {
	op := &request.Operation{
		Name:       opGetAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAttributesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GetAttributesOutput{}
	req.Data = output
	return
}

// Returns all of the attributes associated with the specified item. Optionally,
// the attributes returned can be limited to one or more attributes by specifying
// an attribute name parameter.
//
//  If the item does not exist on the replica that was accessed for this operation,
// an empty set is returned. The system does not return an error as it cannot
// guarantee the item does not exist on other replicas.
//
//  If GetAttributes is called without being passed any attribute names, all
// the attributes for the item are returned.
func (c *SimpleDB) GetAttributes(input *GetAttributesInput) (*GetAttributesOutput, error) {
	req, out := c.GetAttributesRequest(input)
	err := req.Send()
	return out, err
}

const opListDomains = "ListDomains"

// ListDomainsRequest generates a request for the ListDomains operation.
func (c *SimpleDB) ListDomainsRequest(input *ListDomainsInput) (req *request.Request, output *ListDomainsOutput) {
	op := &request.Operation{
		Name:       opListDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxNumberOfDomains",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDomainsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListDomainsOutput{}
	req.Data = output
	return
}

// The ListDomains operation lists all domains associated with the Access Key
// ID. It returns domain names up to the limit set by MaxNumberOfDomains (#MaxNumberOfDomains).
// A NextToken (#NextToken) is returned if there are more than MaxNumberOfDomains
// domains. Calling ListDomains successive times with the NextToken provided
// by the operation returns up to MaxNumberOfDomains more domain names with
// each successive operation call.
func (c *SimpleDB) ListDomains(input *ListDomainsInput) (*ListDomainsOutput, error) {
	req, out := c.ListDomainsRequest(input)
	err := req.Send()
	return out, err
}

func (c *SimpleDB) ListDomainsPages(input *ListDomainsInput, fn func(p *ListDomainsOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.ListDomainsRequest(input)
	page.Handlers.Build.PushBack(request.MakeAddToUserAgentFreeFormHandler("Paginator"))
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*ListDomainsOutput), lastPage)
	})
}

const opPutAttributes = "PutAttributes"

// PutAttributesRequest generates a request for the PutAttributes operation.
func (c *SimpleDB) PutAttributesRequest(input *PutAttributesInput) (req *request.Request, output *PutAttributesOutput) {
	op := &request.Operation{
		Name:       opPutAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutAttributesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &PutAttributesOutput{}
	req.Data = output
	return
}

// The PutAttributes operation creates or replaces attributes in an item. The
// client may specify new attributes using a combination of the Attribute.X.Name
// and Attribute.X.Value parameters. The client specifies the first attribute
// by the parameters Attribute.0.Name and Attribute.0.Value, the second attribute
// by the parameters Attribute.1.Name and Attribute.1.Value, and so on.
//
//  Attributes are uniquely identified in an item by their name/value combination.
// For example, a single item can have the attributes { "first_name", "first_value"
// } and { "first_name", second_value" }. However, it cannot have two attribute
// instances where both the Attribute.X.Name and Attribute.X.Value are the same.
//
//  Optionally, the requestor can supply the Replace parameter for each individual
// attribute. Setting this value to true causes the new attribute value to replace
// the existing attribute value(s). For example, if an item has the attributes
// { 'a', '1' }, { 'b', '2'} and { 'b', '3' } and the requestor calls PutAttributes
// using the attributes { 'b', '4' } with the Replace parameter set to true,
// the final attributes of the item are changed to { 'a', '1' } and { 'b', '4'
// }, which replaces the previous values of the 'b' attribute with the new value.
//
//  Using PutAttributes to replace attribute values that do not exist will
// not result in an error response.   You cannot specify an empty string as
// an attribute name.
//
//  Because Amazon SimpleDB makes multiple copies of client data and uses an
// eventual consistency update model, an immediate GetAttributes or Select operation
// (read) immediately after a PutAttributes or DeleteAttributes operation (write)
// might not return the updated data.
//
//  The following limitations are enforced for this operation:  256 total attribute
// name-value pairs per item One billion attributes per domain 10 GB of total
// user data storage per domain
func (c *SimpleDB) PutAttributes(input *PutAttributesInput) (*PutAttributesOutput, error) {
	req, out := c.PutAttributesRequest(input)
	err := req.Send()
	return out, err
}

const opSelect = "Select"

// SelectRequest generates a request for the Select operation.
func (c *SimpleDB) SelectRequest(input *SelectInput) (req *request.Request, output *SelectOutput) {
	op := &request.Operation{
		Name:       opSelect,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &SelectInput{}
	}

	req = c.newRequest(op, input, output)
	output = &SelectOutput{}
	req.Data = output
	return
}

// The Select operation returns a set of attributes for ItemNames that match
// the select expression. Select is similar to the standard SQL SELECT statement.
//
//  The total size of the response cannot exceed 1 MB in total size. Amazon
// SimpleDB automatically adjusts the number of items returned per page to enforce
// this limit. For example, if the client asks to retrieve 2500 items, but each
// individual item is 10 kB in size, the system returns 100 items and an appropriate
// NextToken so the client can access the next page of results.
//
//  For information on how to construct select expressions, see Using Select
// to Create Amazon SimpleDB Queries in the Developer Guide.
func (c *SimpleDB) Select(input *SelectInput) (*SelectOutput, error) {
	req, out := c.SelectRequest(input)
	err := req.Send()
	return out, err
}

func (c *SimpleDB) SelectPages(input *SelectInput, fn func(p *SelectOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.SelectRequest(input)
	page.Handlers.Build.PushBack(request.MakeAddToUserAgentFreeFormHandler("Paginator"))
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*SelectOutput), lastPage)
	})
}

type Attribute struct {
	_ struct{} `type:"structure"`

	AlternateNameEncoding *string `type:"string"`

	AlternateValueEncoding *string `type:"string"`

	// The name of the attribute.
	Name *string `type:"string" required:"true"`

	// The value of the attribute.
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Attribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Attribute) GoString() string {
	return s.String()
}

type BatchDeleteAttributesInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which the attributes are being deleted.
	DomainName *string `type:"string" required:"true"`

	// A list of items on which to perform the operation.
	Items []*DeletableItem `locationNameList:"Item" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s BatchDeleteAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchDeleteAttributesInput) GoString() string {
	return s.String()
}

type BatchDeleteAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s BatchDeleteAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchDeleteAttributesOutput) GoString() string {
	return s.String()
}

type BatchPutAttributesInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which the attributes are being stored.
	DomainName *string `type:"string" required:"true"`

	// A list of items on which to perform the operation.
	Items []*ReplaceableItem `locationNameList:"Item" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s BatchPutAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchPutAttributesInput) GoString() string {
	return s.String()
}

type BatchPutAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s BatchPutAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchPutAttributesOutput) GoString() string {
	return s.String()
}

type CreateDomainInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain to create. The name can range between 3 and 255 characters
	// and can contain the following characters: a-z, A-Z, 0-9, '_', '-', and '.'.
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDomainInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDomainInput) GoString() string {
	return s.String()
}

type CreateDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDomainOutput) GoString() string {
	return s.String()
}

type DeletableAttribute struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	Name *string `type:"string" required:"true"`

	// The value of the attribute.
	Value *string `type:"string"`
}

// String returns the string representation
func (s DeletableAttribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletableAttribute) GoString() string {
	return s.String()
}

type DeletableItem struct {
	_ struct{} `type:"structure"`

	Attributes []*DeletableAttribute `locationNameList:"Attribute" type:"list" flattened:"true"`

	Name *string `locationName:"ItemName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletableItem) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletableItem) GoString() string {
	return s.String()
}

type DeleteAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of Attributes. Similar to columns on a spreadsheet, attributes represent
	// categories of data that can be assigned to items.
	Attributes []*DeletableAttribute `locationNameList:"Attribute" type:"list" flattened:"true"`

	// The name of the domain in which to perform the operation.
	DomainName *string `type:"string" required:"true"`

	// The update condition which, if specified, determines whether the specified
	// attributes will be deleted or not. The update condition must be satisfied
	// in order for this request to be processed and the attributes to be deleted.
	Expected *UpdateCondition `type:"structure"`

	// The name of the item. Similar to rows on a spreadsheet, items represent individual
	// objects that contain one or more value-attribute pairs.
	ItemName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAttributesInput) GoString() string {
	return s.String()
}

type DeleteAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAttributesOutput) GoString() string {
	return s.String()
}

type DeleteDomainInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain to delete.
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDomainInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDomainInput) GoString() string {
	return s.String()
}

type DeleteDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDomainOutput) GoString() string {
	return s.String()
}

type DomainMetadataInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain for which to display the metadata of.
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DomainMetadataInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainMetadataInput) GoString() string {
	return s.String()
}

type DomainMetadataOutput struct {
	_ struct{} `type:"structure"`

	// The number of unique attribute names in the domain.
	AttributeNameCount *int64 `type:"integer"`

	// The total size of all unique attribute names in the domain, in bytes.
	AttributeNamesSizeBytes *int64 `type:"long"`

	// The number of all attribute name/value pairs in the domain.
	AttributeValueCount *int64 `type:"integer"`

	// The total size of all attribute values in the domain, in bytes.
	AttributeValuesSizeBytes *int64 `type:"long"`

	// The number of all items in the domain.
	ItemCount *int64 `type:"integer"`

	// The total size of all item names in the domain, in bytes.
	ItemNamesSizeBytes *int64 `type:"long"`

	// The data and time when metadata was calculated, in Epoch (UNIX) seconds.
	Timestamp *int64 `type:"integer"`
}

// String returns the string representation
func (s DomainMetadataOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainMetadataOutput) GoString() string {
	return s.String()
}

type GetAttributesInput struct {
	_ struct{} `type:"structure"`

	// The names of the attributes.
	AttributeNames []*string `locationNameList:"AttributeName" type:"list" flattened:"true"`

	// Determines whether or not strong consistency should be enforced when data
	// is read from SimpleDB. If true, any data previously written to SimpleDB will
	// be returned. Otherwise, results will be consistent eventually, and the client
	// may not see data that was written immediately before your read.
	ConsistentRead *bool `type:"boolean"`

	// The name of the domain in which to perform the operation.
	DomainName *string `type:"string" required:"true"`

	// The name of the item.
	ItemName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAttributesInput) GoString() string {
	return s.String()
}

type GetAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The list of attributes returned by the operation.
	Attributes []*Attribute `locationNameList:"Attribute" type:"list" flattened:"true"`
}

// String returns the string representation
func (s GetAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAttributesOutput) GoString() string {
	return s.String()
}

type Item struct {
	_ struct{} `type:"structure"`

	AlternateNameEncoding *string `type:"string"`

	// A list of attributes.
	Attributes []*Attribute `locationNameList:"Attribute" type:"list" flattened:"true" required:"true"`

	// The name of the item.
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Item) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Item) GoString() string {
	return s.String()
}

type ListDomainsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of domain names you want returned. The range is 1 to 100.
	// The default setting is 100.
	MaxNumberOfDomains *int64 `type:"integer"`

	// A string informing Amazon SimpleDB where to start the next list of domain
	// names.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDomainsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDomainsInput) GoString() string {
	return s.String()
}

type ListDomainsOutput struct {
	_ struct{} `type:"structure"`

	// A list of domain names that match the expression.
	DomainNames []*string `locationNameList:"DomainName" type:"list" flattened:"true"`

	// An opaque token indicating that there are more domains than the specified
	// MaxNumberOfDomains still available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDomainsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDomainsOutput) GoString() string {
	return s.String()
}

type PutAttributesInput struct {
	_ struct{} `type:"structure"`

	// The list of attributes.
	Attributes []*ReplaceableAttribute `locationNameList:"Attribute" type:"list" flattened:"true" required:"true"`

	// The name of the domain in which to perform the operation.
	DomainName *string `type:"string" required:"true"`

	// The update condition which, if specified, determines whether the specified
	// attributes will be updated or not. The update condition must be satisfied
	// in order for this request to be processed and the attributes to be updated.
	Expected *UpdateCondition `type:"structure"`

	// The name of the item.
	ItemName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutAttributesInput) GoString() string {
	return s.String()
}

type PutAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutAttributesOutput) GoString() string {
	return s.String()
}

type ReplaceableAttribute struct {
	_ struct{} `type:"structure"`

	// The name of the replaceable attribute.
	Name *string `type:"string" required:"true"`

	// A flag specifying whether or not to replace the attribute/value pair or to
	// add a new attribute/value pair. The default setting is false.
	Replace *bool `type:"boolean"`

	// The value of the replaceable attribute.
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceableAttribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ReplaceableAttribute) GoString() string {
	return s.String()
}

type ReplaceableItem struct {
	_ struct{} `type:"structure"`

	// The list of attributes for a replaceable item.
	Attributes []*ReplaceableAttribute `locationNameList:"Attribute" type:"list" flattened:"true" required:"true"`

	// The name of the replaceable item.
	Name *string `locationName:"ItemName" type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceableItem) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ReplaceableItem) GoString() string {
	return s.String()
}

type SelectInput struct {
	_ struct{} `type:"structure"`

	// Determines whether or not strong consistency should be enforced when data
	// is read from SimpleDB. If true, any data previously written to SimpleDB will
	// be returned. Otherwise, results will be consistent eventually, and the client
	// may not see data that was written immediately before your read.
	ConsistentRead *bool `type:"boolean"`

	// A string informing Amazon SimpleDB where to start the next list of ItemNames.
	NextToken *string `type:"string"`

	// The expression used to query the domain.
	SelectExpression *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SelectInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SelectInput) GoString() string {
	return s.String()
}

type SelectOutput struct {
	_ struct{} `type:"structure"`

	// A list of items that match the select expression.
	Items []*Item `locationNameList:"Item" type:"list" flattened:"true"`

	// An opaque token indicating that more items than MaxNumberOfItems were matched,
	// the response size exceeded 1 megabyte, or the execution time exceeded 5 seconds.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s SelectOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SelectOutput) GoString() string {
	return s.String()
}

// Specifies the conditions under which data should be updated. If an update
// condition is specified for a request, the data will only be updated if the
// condition is satisfied. For example, if an attribute with a specific name
// and value exists, or if a specific attribute doesn't exist.
type UpdateCondition struct {
	_ struct{} `type:"structure"`

	// A value specifying whether or not the specified attribute must exist with
	// the specified value in order for the update condition to be satisfied. Specify
	// true if the attribute must exist for the update condition to be satisfied.
	// Specify false if the attribute should not exist in order for the update condition
	// to be satisfied.
	Exists *bool `type:"boolean"`

	// The name of the attribute involved in the condition.
	Name *string `type:"string"`

	// The value of an attribute. This value can only be specified when the Exists
	// parameter is equal to true.
	Value *string `type:"string"`
}

// String returns the string representation
func (s UpdateCondition) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateCondition) GoString() string {
	return s.String()
}
