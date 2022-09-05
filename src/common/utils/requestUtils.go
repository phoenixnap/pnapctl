package utils

import "net/http"

// Utility functions

// Extracts the ID (first argument) passed in the arguments.
// If no ID was found, returns nil instead.
func ExtractId(args []string) *string {
	if len(args) == 1 {
		return &args[0]
	}
	return nil
}

// Request builder pattern

// Struct used for requests that utilize an ID.
type IdRequestHandler[T any] struct {
	id          *string
	request     Request[*T]
	requestList Request[[]T]
	print       Printer[*T]
	printList   Printer[[]T]
}

// Represents a normal request that takes no parameters.
type Req[T any] func() (T, *http.Response, error)

// Represents a normal request that takes an extra parameter.
// Examples include a request body, query params, or even an ID.
type ReqParam[T, P any] func(P) (T, *http.Response, error)

// Represents a response - wraps its data in order to add
// functionality.
type Response[T any] struct {
	body        T
	err         error
	commandName string
}

// Represents a request - wraps the function in order to
// add more functionality.
type Request[T any] struct {
	request Req[T]
}

// Represents a 'printer' method
type Printer[T any] func(T, string) error

// Represents a 'printer' method, including the 'full' boolean.
type PrinterWithFull[T any] func(T, bool, string) error

// ID specific Logic

// Parses the ID passed in order and stores the result in a struct.
// 'T' should be set to the type that the eventual request will return.
//
// Used to setup two requests - using one depending on whether an ID
// was passed, for example GetById vs. GetAll
//
//	// Prepare request for *get server*
//	UseIdFor[bmcapisdk.Server](args)
func UseIdFor[T any](args []string) IdRequestHandler[T] {
	return IdRequestHandler[T]{
		id: ExtractId(args),
	}
}

// Prepares the request and printer to use 'if' the ID was passed.
//
//	UseIdFor[bmcapisdk.Server](args).
//		IfPresent(
//			bmcapi.Client.ServersGetById,
//			utils.UsingFull(Full, printer.PrintServersResponse)
//		)
func (p IdRequestHandler[T]) IfPresent(requestWith ReqParam[*T, string], printWith Printer[*T]) IdRequestHandler[T] {
	if p.id != nil {
		p.request = DoRequestWith(requestWith, *p.id)
		p.print = printWith
	}
	return p
}

// Prepares the request and printer to use 'if' the ID was not passed.
//
//	UseIdFor[bmcapisdk.Server](args).
//		IfPresent(
//			bmcapi.Client.ServersGet,
//			utils.UsingFull(Full, printer.PrintServerResponse)
//		).
//		Else(
//			utils.DoRequestWith(bmcapi.Client.ServersGetById, *queryParams),
//			utils.UsingFull(Full, printer.PrintServerListResponse)
//		)
func (p IdRequestHandler[T]) Else(requestWith Request[[]T], printWith Printer[[]T]) IdRequestHandler[T] {
	p.requestList = requestWith
	p.printList = printWith
	return p
}

// Executes the configured request. Returns any errors that may have occured.
//
//	UseIdFor[bmcapisdk.Server](args).
//		IfPresent(
//			bmcapi.Client.ServersGet,
//			utils.UsingFull(Full, printer.PrintServerResponse)
//		).
//		Else(
//			utils.DoRequestWith(bmcapi.Client.ServersGetById, *queryParams),
//			utils.UsingFull(Full, printer.PrintServerListResponse)
//		).
//		Execute("get servers")
func (p IdRequestHandler[T]) Execute(commandName string) error {
	if p.id != nil {
		return p.request.Execute(commandName).ThenPrint(p.print)
	}
	return p.requestList.Execute(commandName).ThenPrint(p.printList)
}

// Request Logic

// Prepares a request for execution.
func DoRequest[T any](request Req[T]) Request[T] {
	return Request[T]{
		request: request,
	}
}

// Prepares a request with the parameter passed for execution.
func DoRequestWith[T, P any](request ReqParam[T, P], param P) Request[T] {
	return DoRequest(func() (T, *http.Response, error) {
		return request(param)
	})
}

// Response Logic

// Executes a request, returning a Response[T] object that wraps the actual response.
func (req Request[T]) Execute(commandName string) Response[T] {
	body, resp, err := req.request()

	generatedError := CheckForErrors(resp, err, commandName)

	return Response[T]{
		body:        body,
		err:         *generatedError,
		commandName: commandName,
	}
}

// Printer Logic

// Curries a 'printer' method with a 'full' parameter into a normal Printer[T]
func UsingFull[T any](full bool, printer PrinterWithFull[T]) Printer[T] {
	return func(body T, commandName string) error {
		return printer(body, full, commandName)
	}
}

// Attempts to print the response.
//
// If the response was in error - nothing will be printed and the error will
// be returned immediately.
//
//	response.ThenPrint(printer.PrintTagListResponse)
//
//	// For printers that use 'full', you do this...
//	response.ThenPrint(utils.UsingFull(Full, printer.PrintServerListResponse))
func (resp Response[T]) ThenPrint(printer Printer[T]) error {
	if resp.err != nil {
		return resp.err
	}

	return printer(resp.body, resp.commandName)
}
