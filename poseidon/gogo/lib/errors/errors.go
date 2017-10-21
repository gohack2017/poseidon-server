package errors

var (
	AccessDenied                = Error{403, "AccessDenied", "Access Denied"}
	AuthFailure                 = Error{403, "AuthFailure", "The provided credentials could not be validated. You may not be authorized to carry out the request; Ensure that your account is authorized to use the service and that you are using the correct access keys."}
	Blocked                     = Error{403, "Blocked", "Your account is currently blocked. Contact us if you have questions."}
	IdempotentParameterMismatch = Error{400, "IdempotentParameterMismatch", "The request uses the same client token as a previous, but non-identical request. Do not reuse a client token with different requests, unless the requests are identical."}
	IncompleteSignature         = Error{400, "IncompleteSignature", "The request signature does not conform to service standards."}
	InvalidAction               = Error{400, "InvalidAction", "The action or operation requested is not valid. Verify that the action is typed correctly."}
	InvalidClientTokenID        = Error{400, "InvalidClientTokenID", "The X.509 certificate or access key ID provided does not exist in our records."}
	InvalidPaginationToken      = Error{400, "InvalidPaginationToken", "The specified pagination token is not valid or is expired."}
	InvalidParameter            = Error{400, "InvalidParameter", "A parameter specified in a request is not valid, is unsupported, or cannot be used."}
	InvalidParameterCombination = Error{400, "InvalidParameterCombination", "Indicates an incorrect combination of parameters, or a missing parameter."}
	InvalidParameterValue       = Error{400, "InvalidParameterValue", "A value specified in a parameter is not valid, is unsupported, or cannot be used. Ensure that you specify a resource by using its full ID."}
	InvalidQueryParameter       = Error{400, "InvalidQueryParameter", "The query string is malformed or does not adhere to service standards."}
	MalformedParameter          = Error{400, "MalformedParameter", "The parameter specified in a request is not valid, is contains a syntax error, or cannot be decoded."}
	MalformedQueryString        = Error{400, "MalformedQueryString", "The query string contains a syntax error."}
	MissingAction               = Error{400, "MissingAction", "The request is missing an action or a required parameter."}
	MissingAuthenticationToken  = Error{400, "MissingAuthenticationToken", "The request must contain either a valid (registered) access key ID or X.509 certificate."}
	MissingParameter            = Error{400, "MissingParameter", "The request is missing a required parameter. Ensure that you have supplied all the required parameters for the request."}
	MissingSecurityElement      = Error{400, "MissingSecurityElement", "The request is missing a security element."}
	MissingSecurityHeader       = Error{400, "MissingSecurityHeader", "Your request is missing a required header."}
	PendingVerification         = Error{400, "PendingVerification", "Your account is pending verification. Until the verification process is complete, you may not be able to carry out requests with this account. If you have questions, contact us Support."}
	RequestExpired              = Error{400, "RequestExpired", "The request reached the service more than 15 minutes after the date stamp on the request or more than 15 minutes after the request expiration date (such as for pre-signed URLs), or the date stamp on the request is more than 15 minutes in the future. If you're using temporary security credentials, this error can also occur if the credentials have expired."}
	UnauthorizedOperation       = Error{400, "UnauthorizedOperation", "You are not authorized to perform this operation."}
	UnknownParameter            = Error{400, "UnknownParameter", "An unknown or unrecognized parameter was supplied. Requests that could cause this error include supplying a misspelled parameter or a parameter that is not supported for the specified API version."}
	UnsupportedProtocol         = Error{400, "UnsupportedProtocol", "SOAP has been deprecated and is no longer supported."}
	ValidationError             = Error{400, "ValidationError", "The input fails to satisfy the constraints specified by an service."}
	InvalidServiceName          = Error{400, "InvalidServiceName", "The name of the service is not valid."}
	Unsupported                 = Error{400, "Unsupported", "The specified request is unsupported."}
	UnsupportedOperation        = Error{400, "UnsupportedOperation", "The specified request includes an unsupported operation."}
	InternalError               = Error{500, "InternalError", "An internal error has occurred. Retry your request, but if the problem persists, contact us with details by posting a message on the service forums."}
	InternalFailure             = Error{400, "InternalFailure", "The request processing has failed because of an unknown error, exception or failure."}
	RequestLimitExceeded        = Error{400, "RequestLimitExceeded", "The maximum request rate permitted by the APIs has been exceeded for your account. For best results, use an increasing or variable sleep interval between requests."}
	ServiceUnavailable          = Error{400, "ServiceUnavailable", "The request has failed due to a temporary failure of the server."}
	SignatureDoesNotMatch       = Error{403, "SignatureDoesNotMatch", "The request signature we calculated does not match the signature you provided."}
	Unavailable                 = Error{400, "Unavailable", "The server is overloaded and can't handle the request."}
	PoliceAlreadyExists         = Error{409, "PoliceAlreadyExists", "The requested police station info is not available."}

	HotelAlreadyExists = Error{409, "HotelAlreadyExists", "The requested hotel info is not available."}
)
