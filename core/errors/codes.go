package errors

import (
	"errors"
	"net/http"
)

// Codes of error
type Codes uint8

// error list
const (
	// Unclassified error.
	Other Codes = iota
	// unsupported input type
	UnsupportedDataType
	// Database type is not exists
	DatabaseTypeNotExists
	// No data found error
	NoDataFound
	// Data exists
	DataExists
	//Database execution fail
	DatabaseExecutionFail
	//Database Tx Fail
	DatabaseTxFail
	// error because redis connection is not exists
	RedisNotExists
	// error duplicate because SetNX
	RedisKeyDuplicate
	// error because redis key not found
	RedisKeyNotFound
	//Redis Operation Failed
	RedisOperationFail
	//Mq Publish Failed
	MQPublishFail
	// invalid request
	InvalidRequest
	//invalid state machine request
	InvalidStateMachineRequest
	// apicalls
	// cannot api call
	APICallsFail
	// can call api but got error
	APISuccessError
	// InvalidFormatField is used when we failed to parse field which does not follow the rule
	InvalidFormatField
	// MissingRequiredField
	MissingRequiredField
	//MissingIdempotencyKey
	MissingIdempotencyKey
	//InvalidIdempotencyKey
	InvalidIdempotencyKey
	//InvalidAccessToken
	InvalidAccessToken
	//InvalidClient
	InvalidClientUsername
	//Internal Server Error
	InternalServerError
	// FailedMarshalOrUnmarshal json to/from object
	FailedMarshalOrUnmarshal
	// FailedMarshalOrUnmarshal json to/from object
	MQPublisherNoInstance
	//State up to date
	StateUpToDate
	//Invalid user id
	InvalidUserID
	//Invalida user name
	InvalidUsername
	//State machine not found in memory
	StateMachineMemoryNotFound
	//State machine not found
	StateMachineNotFound
	//Invalid Business ID
	InvalidBusinessID
	//InvalidBusinessCode
	InvalidBusinessCode
	//Internal Request API Fail
	InternalRequestAPIFailed
	//Invalid Tx V2 Request
	InvalidTxV2OperationRequest
	//FsmOperationError
	FsmOperationError
	//FsmReinitializeError
	FsmReinitializeError
	//InvalidOriginState
	InvalidOriginState
	//InvalidStateTransition
	InvalidStateTransition
	//CreateTokenFailure
	CreateTokenFailure
	//InvalidWebhookRequest
	InvalidWebHookRequest
	//WebhookNilRequest
	WebHookNilRequest
	//WebhookEmptyOrder
	WebHookEmptyOrder
	//WebHookEmptyItem
	WebHookEmptyItem
	//WebHookUnsupportedBusinessCode
	WebHookUnsupportedBusinessCode
	//WebHookUnsupportedOrderStatus
	WebHookUnsupportedOrderStatus
	//WebHookInvalidServerConf
	WebHookInvalidServerConf
	//TooManyRequest
	TooManyRequest
)

//Error string
const (
	RedisKeyNotFoundErr         = "Redis key is not found"
	FailMarshallOrUnMarshallErr = "Failed marshal or unmarshal"
)

// GetErrorAndCode func
func (c Codes) GetErrorAndCode() (string, int) {
	switch c {
	case Other:
		return "Internal server error", http.StatusInternalServerError
	case FailedMarshalOrUnmarshal:
		return "Failed marshal or unmarshal", http.StatusInternalServerError
	case DatabaseTypeNotExists:
		return "Database type is not exists", http.StatusInternalServerError
	case DatabaseExecutionFail:
		return "Error executing database operation", http.StatusInternalServerError
	case DatabaseTxFail:
		return "Cannot get database transaction", http.StatusInternalServerError
	case DataExists:
		return "Data exists", http.StatusBadRequest
	case NoDataFound:
		return "No data found", http.StatusBadRequest
	case MQPublisherNoInstance:
		return "MQ Publisher need to init first", http.StatusInternalServerError
	case InvalidFormatField:
		return "Failed to parse. Invalid field format.", http.StatusBadRequest
	case RedisOperationFail:
		return "Cannot execute Redis operation", http.StatusInternalServerError
	case StateUpToDate:
		return "Status is already up to date", http.StatusBadRequest
	case InvalidClientUsername:
		return "Invalid client username", http.StatusForbidden
	case InvalidUserID:
		return "Invalid used id", http.StatusForbidden
	case InvalidUsername:
		return "Invalid username", http.StatusForbidden
	case StateMachineMemoryNotFound:
		return "State machine not found in memory", http.StatusInternalServerError
	case StateMachineNotFound:
		return "State machine not found", http.StatusBadRequest
	case InvalidRequest:
		return "Invalid request", http.StatusBadRequest
	case InvalidAccessToken:
		return "Invalid access token", http.StatusForbidden
	case InvalidBusinessID:
		return "Invalid business ID", http.StatusForbidden
	case InvalidBusinessCode:
		return "Invalid business code", http.StatusForbidden
	case InternalRequestAPIFailed:
		return "Request API to internal service failed", http.StatusInternalServerError
	case InvalidTxV2OperationRequest:
		return "Could not determine Tx V2 operation request", http.StatusBadRequest
	case FsmOperationError:
		return "Client state operation error", http.StatusInternalServerError
	case InvalidStateMachineRequest:
		return "Invalid state machine request", http.StatusBadRequest
	case CreateTokenFailure:
		return "Fail on creating the token", http.StatusInternalServerError
	case FsmReinitializeError:
		return "Fail on reinitializing the Fsm", http.StatusInternalServerError
	case InvalidWebHookRequest:
		return "Invalid web-hook request", http.StatusInternalServerError
	case MissingIdempotencyKey:
		return "Missing Idempotency-Key", http.StatusBadRequest
	case InvalidIdempotencyKey:
		return "Invalid Idempotency-Key", http.StatusBadRequest
	case RedisKeyNotFound:
		return RedisKeyNotFoundErr, http.StatusBadRequest
	case InvalidOriginState:
		return "Invalid origin state", http.StatusBadRequest
	case InvalidStateTransition:
		return "Invalid state transition", http.StatusBadRequest
	case MQPublishFail:
		return "Cannot publish data to MQ", http.StatusInternalServerError
	case WebHookNilRequest:
		return "Nil Webhook Request", http.StatusBadRequest
	case WebHookEmptyOrder:
		return "Webhook Empty Order", http.StatusBadRequest
	case WebHookEmptyItem:
		return "Webhook Empty Item", http.StatusBadRequest
	case WebHookUnsupportedBusinessCode:
		return "Webhook Unsupported Business Code", http.StatusBadRequest
	case WebHookInvalidServerConf:
		return "Webhook Invalid Server Configuration", http.StatusBadRequest
	case WebHookUnsupportedOrderStatus:
		return "Webhook Unsupported Order Status", http.StatusBadRequest
	case TooManyRequest:
		return "Temporary rejected due exceeding request number", http.StatusTooManyRequests
	default:
		return "Internal server error", http.StatusInternalServerError
	}
}

// Err return standard error type
func (c Codes) Err() error {
	errString, _ := c.GetErrorAndCode()
	return errors.New(errString)
}

func (c Codes) GetErrAndCode() (error, int) {
	errString, code := c.GetErrorAndCode()
	return errors.New(errString), code
}
