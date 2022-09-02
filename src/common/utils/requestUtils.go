package utils

import "net/http"

// Utility functions

func ExtractId(args []string) *string {
	if len(args) == 1 {
		return &args[0]
	}
	return nil
}

// Request builder pattern

// Types for clarity.
type IdParser[T any] struct {
	id          *string
	request     Request[*T]
	requestList Request[[]T]
	print       Printer[*T]
	printList   Printer[[]T]
}

type Req[T any] func() (T, *http.Response, error)
type ReqParam[T, P any] func(P) (T, *http.Response, error)
type Response[T any] struct {
	body        T
	err         error
	commandName string
}

type Request[T any] struct {
	request Req[T]
}

type Printer[T any] func(T, string) error
type PrinterWithFull[T any] func(T, bool, string) error

// ID specific Logic
func UseIdFor[T any](args []string) IdParser[T] {
	var id *string
	if len(args) == 1 {
		id = &args[0]
	}
	return IdParser[T]{
		id: id,
	}
}

func (p IdParser[T]) IfPresent(requestWith ReqParam[*T, string], printWith Printer[*T]) IdParser[T] {
	if p.id != nil {
		p.request = DoRequestWith(requestWith, *p.id)
		p.print = printWith
	}
	return p
}

func (p IdParser[T]) Else(requestWith Request[[]T], printWith Printer[[]T]) IdParser[T] {
	p.requestList = requestWith
	p.printList = printWith
	return p
}

func (p IdParser[T]) Execute(commandName string) error {
	if p.id != nil {
		return p.request.Execute(commandName).ThenPrint(p.print)
	}
	return p.requestList.Execute(commandName).ThenPrint(p.printList)
}

// Request Logic

func DoRequest[T any](request Req[T]) Request[T] {
	return Request[T]{
		request: request,
	}
}

func DoRequestWith[T, P any](request ReqParam[T, P], param P) Request[T] {
	return DoRequest(func() (T, *http.Response, error) {
		return request(param)
	})
}

// Response Logic

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

func UsingFull[T any](full bool, printer PrinterWithFull[T]) Printer[T] {
	return func(body T, commandName string) error {
		return printer(body, full, commandName)
	}
}

func (resp Response[T]) ThenPrint(printer Printer[T]) error {
	if resp.err != nil {
		return resp.err
	}

	return printer(resp.body, resp.commandName)
}
