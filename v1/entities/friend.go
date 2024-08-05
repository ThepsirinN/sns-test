package entities

type CreateFriendRequestRequest struct {
	SourceId int32 `validate:"required"`
	DestId   int32 `json:"dest_id" validate:"required"`
}

type (
	GetAllFriendRequestRequest struct {
		UserId int32 `validate:"required"`
	}

	GetAllFriendRequestResponse struct {
		Id               int32  `json:"id"`
		SourceId         int32  `json:"source_id"`
		DestId           int32  `json:"dest_id"`
		Status           string `json:"status"`
		SourceEmail      string `json:"source_email"`
		SourceFirstName  string `json:"source_first_name"`
		SourceLastName   string `json:"source_last_name"`
		SourceProfileImg string `json:"source_profile_image"`
		DestEmail        string `json:"dest_email"`
		DestFirstName    string `json:"dest_first_name"`
		DestLastName     string `json:"dest_last_name"`
		DestProfileImg   string `json:"dest_profile_image"`
	}
)

type (
	ListFriendRequest struct {
		UserId int32 `validate:"required"`
	}

	ListFriendQuery struct {
		Id               int32  `json:"id"`
		SourceId         int32  `json:"source_id"`
		DestId           int32  `json:"dest_id"`
		SourceEmail      string `json:"source_email"`
		SourceFirstName  string `json:"source_first_name"`
		SourceLastName   string `json:"source_last_name"`
		SourceProfileImg string `json:"source_profile_image"`
		DestEmail        string `json:"dest_email"`
		DestFirstName    string `json:"dest_first_name"`
		DestLastName     string `json:"dest_last_name"`
		DestProfileImg   string `json:"dest_profile_image"`
	}

	ListFriendResponse struct {
		Id               int32  `json:"id"`
		FriendUserID     int32  `json:"friend_user_id"`
		FriendEmail      string `json:"friend_email"`
		FriendFirstName  string `json:"friend_first_name"`
		FriendLastName   string `json:"friend_last_name"`
		FriendProfileImg string `json:"friend_Profile_img"`
	}
)

type UpdateFriendRequestStatusRequest struct {
	Id       int32 `json:"id" validate:"required"`
	SourceId int32 `json:"source_id" validate:"required"`
	DestId   int32 `json:"dest_id" validate:"required"`
}

type DeleteFriendRequestRequest struct {
	Id       int32 `json:"id" validate:"required"`
	SourceId int32 `json:"source_id" validate:"required"`
	DestId   int32 `json:"dest_id" validate:"required"`
}
