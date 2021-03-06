// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCapacityReservationsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Capacity Reservation.
	CapacityReservationIds []string `locationName:"CapacityReservationId" locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the returned
	// nextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeCapacityReservationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCapacityReservationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCapacityReservationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCapacityReservationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Capacity Reservations.
	CapacityReservations []CapacityReservation `locationName:"capacityReservationSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeCapacityReservationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCapacityReservations = "DescribeCapacityReservations"

// DescribeCapacityReservationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your Capacity Reservations. The results describe
// only the Capacity Reservations in the AWS Region that you're currently using.
//
//    // Example sending a request using DescribeCapacityReservationsRequest.
//    req := client.DescribeCapacityReservationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeCapacityReservations
func (c *Client) DescribeCapacityReservationsRequest(input *DescribeCapacityReservationsInput) DescribeCapacityReservationsRequest {
	op := &aws.Operation{
		Name:       opDescribeCapacityReservations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeCapacityReservationsInput{}
	}

	req := c.newRequest(op, input, &DescribeCapacityReservationsOutput{})
	return DescribeCapacityReservationsRequest{Request: req, Input: input, Copy: c.DescribeCapacityReservationsRequest}
}

// DescribeCapacityReservationsRequest is the request type for the
// DescribeCapacityReservations API operation.
type DescribeCapacityReservationsRequest struct {
	*aws.Request
	Input *DescribeCapacityReservationsInput
	Copy  func(*DescribeCapacityReservationsInput) DescribeCapacityReservationsRequest
}

// Send marshals and sends the DescribeCapacityReservations API request.
func (r DescribeCapacityReservationsRequest) Send(ctx context.Context) (*DescribeCapacityReservationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCapacityReservationsResponse{
		DescribeCapacityReservationsOutput: r.Request.Data.(*DescribeCapacityReservationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCapacityReservationsRequestPaginator returns a paginator for DescribeCapacityReservations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCapacityReservationsRequest(input)
//   p := ec2.NewDescribeCapacityReservationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCapacityReservationsPaginator(req DescribeCapacityReservationsRequest) DescribeCapacityReservationsPaginator {
	return DescribeCapacityReservationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCapacityReservationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeCapacityReservationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCapacityReservationsPaginator struct {
	aws.Pager
}

func (p *DescribeCapacityReservationsPaginator) CurrentPage() *DescribeCapacityReservationsOutput {
	return p.Pager.CurrentPage().(*DescribeCapacityReservationsOutput)
}

// DescribeCapacityReservationsResponse is the response type for the
// DescribeCapacityReservations API operation.
type DescribeCapacityReservationsResponse struct {
	*DescribeCapacityReservationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCapacityReservations request.
func (r *DescribeCapacityReservationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
