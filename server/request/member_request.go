package request

type DeleteMembershipRequest struct {
	MembershipName string `json:"subsName" validate:"required"`
}
