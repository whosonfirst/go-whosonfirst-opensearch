// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or overwrites one or more tags for the specified resource. Tags are
// metadata that you can assign to your automations, documents, managed nodes,
// maintenance windows, Parameter Store parameters, and patch baselines. Tags
// enable you to categorize your resources in different ways, for example, by
// purpose, owner, or environment. Each tag consists of a key and an optional
// value, both of which you define. For example, you could define a set of tags for
// your account's managed nodes that helps you track each node's owner and stack
// level. For example:
//   - Key=Owner,Value=DbAdmin
//   - Key=Owner,Value=SysAdmin
//   - Key=Owner,Value=Dev
//   - Key=Stack,Value=Production
//   - Key=Stack,Value=Pre-Production
//   - Key=Stack,Value=Test
//
// Most resources can have a maximum of 50 tags. Automations can have a maximum of
// 5 tags. We recommend that you devise a set of tag keys that meets your needs for
// each resource type. Using a consistent set of tag keys makes it easier for you
// to manage your resources. You can search and filter the resources based on the
// tags you add. Tags don't have any semantic meaning to and are interpreted
// strictly as a string of characters. For more information about using tags with
// Amazon Elastic Compute Cloud (Amazon EC2) instances, see Tagging your Amazon
// EC2 resources (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html)
// in the Amazon EC2 User Guide.
func (c *Client) AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) {
	if params == nil {
		params = &AddTagsToResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddTagsToResource", params, optFns, c.addOperationAddTagsToResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddTagsToResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddTagsToResourceInput struct {

	// The resource ID you want to tag. Use the ID of the resource. Here are some
	// examples: MaintenanceWindow : mw-012345abcde PatchBaseline : pb-012345abcde
	// Automation : example-c160-4567-8519-012345abcde OpsMetadata object: ResourceID
	// for tagging is created from the Amazon Resource Name (ARN) for the object.
	// Specifically, ResourceID is created from the strings that come after the word
	// opsmetadata in the ARN. For example, an OpsMetadata object with an ARN of
	// arn:aws:ssm:us-east-2:1234567890:opsmetadata/aws/ssm/MyGroup/appmanager has a
	// ResourceID of either aws/ssm/MyGroup/appmanager or /aws/ssm/MyGroup/appmanager .
	// For the Document and Parameter values, use the name of the resource. If you're
	// tagging a shared document, you must use the full ARN of the document.
	// ManagedInstance : mi-012345abcde The ManagedInstance type for this API
	// operation is only for on-premises managed nodes. You must specify the name of
	// the managed node in the following format: mi-ID_number . For example,
	// mi-1a2b3c4d5e6f .
	//
	// This member is required.
	ResourceId *string

	// Specifies the type of resource you are tagging. The ManagedInstance type for
	// this API operation is for on-premises managed nodes. You must specify the name
	// of the managed node in the following format: mi-ID_number . For example,
	// mi-1a2b3c4d5e6f .
	//
	// This member is required.
	ResourceType types.ResourceTypeForTagging

	// One or more tags. The value parameter is required. Don't enter personally
	// identifiable information in this field.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type AddTagsToResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddTagsToResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddTagsToResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTagsToResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addAddTagsToResourceResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAddTagsToResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToResource(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAddTagsToResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "AddTagsToResource",
	}
}

type opAddTagsToResourceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opAddTagsToResourceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opAddTagsToResourceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ssm"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ssm"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ssm")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addAddTagsToResourceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opAddTagsToResourceResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
