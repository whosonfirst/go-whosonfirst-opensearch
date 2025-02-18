// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the status of the specified access key from Active to Inactive, or vice
// versa. This operation can be used to disable a user's key as part of a key
// rotation workflow.
//
// If the UserName is not specified, the user name is determined implicitly based
// on the Amazon Web Services access key ID used to sign the request. If a
// temporary access key is used, then UserName is required. If a long-term key is
// assigned to the user, then UserName is not required. This operation works for
// access keys under the Amazon Web Services account. Consequently, you can use
// this operation to manage Amazon Web Services account root user credentials even
// if the Amazon Web Services account has no associated users.
//
// For information about rotating keys, see [Managing keys and certificates] in the IAM User Guide.
//
// [Managing keys and certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingCredentials.html
func (c *Client) UpdateAccessKey(ctx context.Context, params *UpdateAccessKeyInput, optFns ...func(*Options)) (*UpdateAccessKeyOutput, error) {
	if params == nil {
		params = &UpdateAccessKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccessKey", params, optFns, c.addOperationUpdateAccessKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccessKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccessKeyInput struct {

	// The access key ID of the secret access key you want to update.
	//
	// This parameter allows (through its [regex pattern]) a string of characters that can consist of
	// any upper or lowercased letter or digit.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	AccessKeyId *string

	//  The status you want to assign to the secret access key. Active means that the
	// key can be used for programmatic calls to Amazon Web Services, while Inactive
	// means that the key cannot be used.
	//
	// This member is required.
	Status types.StatusType

	// The name of the user whose key you want to update.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	UserName *string

	noSmithyDocumentSerde
}

type UpdateAccessKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccessKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateAccessKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAccessKey"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateAccessKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccessKey(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateAccessKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAccessKey",
	}
}
