package mappers

import (
	"banks-api/internal/pkg/model"
	pbbanks "banks-api/pkg/banks"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelToPBReview(review *model.Review) *pbbanks.Review {
	if review == nil {
		return nil
	}

	return &pbbanks.Review{
		Id:         review.ID,
		Content:    review.Content,
		IsApproved: review.IsApproved,
		UserEmail:  review.UserEmail,
		UserPhone:  review.UserPhone,
		Rating:     review.Rating,
		BankId:     review.BankID,
		UserName:   review.UserName,
		Data:       timestamppb.New(review.Date),
		Bank:       review.BankName,
	}
}

func ModelsToPBReviews(reviews []*model.Review) []*pbbanks.Review {
	if reviews == nil {
		return nil
	}

	pbReviews := make([]*pbbanks.Review, 0, len(reviews))

	for _, r := range reviews {
		pbReviews = append(pbReviews, ModelToPBReview(r))
	}

	return pbReviews
}

func PBToModelReview(review *pbbanks.Review) *model.Review {
	if review == nil {
		return nil
	}

	return &model.Review{
		ID:         review.Id,
		Content:    review.Content,
		IsApproved: review.IsApproved,
		UserEmail:  review.UserEmail,
		UserPhone:  review.UserPhone,
		Rating:     review.Rating,
		BankID:     review.BankId,
	}
}
