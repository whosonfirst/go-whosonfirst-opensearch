// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the details about a specific task run as part of a maintenance window
// execution.
func (c *Client) GetMaintenanceWindowExecutionTask(ctx context.Context, params *GetMaintenanceWindowExecutionTaskInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionTaskOutput, error) {
	if params == nil {
		params = &GetMaintenanceWindowExecutionTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMaintenanceWindowExecutionTask", params, optFns, c.addOperationGetMaintenanceWindowExecutionTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMaintenanceWindowExecutionTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowExecutionTaskInput struct {

	// The ID of the specific task execution in the maintenance window task that
	// should be retrieved.
	//
	// This member is required.
	TaskId *string

	// The ID of the maintenance window execution that includes the task.
	//
	// This member is required.
	WindowExecutionId *string

	noSmithyDocumentSerde
}

type GetMaintenanceWindowExecutionTaskOutput struct {

	// The details for the CloudWatch alarm you applied to your maintenance window
	// task.
	AlarmConfiguration *types.AlarmConfiguration

	// The time the task execution completed.
	EndTime *time.Time

	// The defined maximum number of task executions that could be run in parallel.
	MaxConcurrency *string

	// The defined maximum number of task execution errors allowed before scheduling
	// of the task execution would have been stopped.
	MaxErrors *string

	// The priority of the task.
	Priority int32

	// The role that was assumed when running the task.
	ServiceRole *string

	// The time the task execution started.
	StartTime *time.Time

	// The status of the task.
	Status types.MaintenanceWindowExecutionStatus

	// The details explaining the status. Not available for all status values.
	StatusDetails *string

	// The Amazon Resource Name (ARN) of the task that ran.
	TaskArn *string

	// The ID of the specific task execution in the maintenance window task that was
	// retrieved.
	TaskExecutionId *string

	// The parameters passed to the task when it was run. TaskParameters has been
	// deprecated. To specify parameters to pass to a task when it runs, instead use
	// the Parameters option in the TaskInvocationParameters structure. For
	// information about how Systems Manager handles these options for the supported
	// maintenance window task types, see MaintenanceWindowTaskInvocationParameters .
	// The map has the following format:
	//   - Key : string, between 1 and 255 characters
	//   - Value : an array of strings, each between 1 and 255 characters
	TaskParameters []map[string]types.MaintenanceWindowTaskParameterValueExpression

	// The CloudWatch alarms that were invoked by the maintenance window task.
	TriggeredAlarms []types.AlarmStateInformation

	// The type of task that was run.
	Type types.MaintenanceWindowTaskType

	// The ID of the maintenance window execution that includes the task.
	WindowExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMaintenanceWindowExecutionTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindowExecutionTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindowExecutionTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMaintenanceWindowExecutionTask"); err != nil {
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
	if err = addOpGetMaintenanceWindowExecutionTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTask(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMaintenanceWindowExecutionTask",
	}
}
