package keeper

import (
	"context"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/empowerchain/empowerchain/x/plasticcredit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ plasticcredit.QueryServer = Querier{}

func (k Querier) Params(goCtx context.Context, req *plasticcredit.QueryParamsRequest) (*plasticcredit.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params, err := k.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	return &plasticcredit.QueryParamsResponse{Params: params}, nil
}

func (k Querier) Issuers(goCtx context.Context, req *plasticcredit.QueryIssuersRequest) (*plasticcredit.QueryIssuersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	issuers, pageRes, err := k.GetIssuers(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &plasticcredit.QueryIssuersResponse{
		Issuers:    issuers,
		Pagination: pageRes,
	}, nil
}

func (k Querier) Issuer(goCtx context.Context, req *plasticcredit.QueryIssuerRequest) (*plasticcredit.QueryIssuerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	issuer, found := k.GetIssuer(ctx, req.IssuerId)
	if !found {
		return nil, errors.Wrapf(plasticcredit.ErrIssuerNotFound, "issuer with id: %d was not found", req.IssuerId)
	}

	return &plasticcredit.QueryIssuerResponse{
		Issuer: &issuer,
	}, nil
}

func (k Querier) Applicant(goCtx context.Context, req *plasticcredit.QueryApplicantRequest) (*plasticcredit.QueryApplicantResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	applicant, found := k.GetApplicant(ctx, req.ApplicantId)
	if !found {
		return nil, errors.Wrapf(plasticcredit.ErrApplicantNotFound, "applicant with id: %d was not found", req.ApplicantId)
	}

	return &plasticcredit.QueryApplicantResponse{
		Applicant: &applicant,
	}, nil
}