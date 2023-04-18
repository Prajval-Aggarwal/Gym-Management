package request

type DeleteMembershipRequest struct {
	MembershipName string `json:"membershipName" validate:"required"`
}
