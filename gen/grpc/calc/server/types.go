// Code generated by goa v3.0.7, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen msa-calculator/design

package server

import (
	calc "msa-calculator/gen/calc"
	calcpb "msa-calculator/gen/grpc/calc/pb"
)

// NewAddPayload builds the payload of the "add" endpoint of the "calc" service
// from the gRPC request type.
func NewAddPayload(message *calcpb.AddRequest) *calc.AddPayload {
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "calc" service.
func NewAddResponse(result int) *calcpb.AddResponse {
	message := &calcpb.AddResponse{}
	message.Field = int32(result)
	return message
}
