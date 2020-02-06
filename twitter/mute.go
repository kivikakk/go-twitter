package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

// MuteService provides methods for mute specific users.
type MuteService struct {
	sling *sling.Sling
}

// newMuteService returns a new MuteService.
func newMuteService(sling *sling.Sling) *MuteService {
	return &MuteService{
		sling: sling.Path("mutes/users/"),
	}
}

// MuteCreateParams are the parameters for MuteService.Create.
type MuteCreateParams struct {
	ScreenName string `url:"screen_name,omitempty,comma"`
	UserID     int64  `url:"user_id,omitempty,comma"`
}

// Create a mute for specific user, return the user muted as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-mutes-users-create
func (s *MuteService) Create(params *MuteCreateParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// MuteDestroyParams are the parameters for MuteService.Destroy.
type MuteDestroyParams struct {
	ScreenName string `url:"screen_name,omitempty,comma"`
	UserID     int64  `url:"user_id,omitempty,comma"`
}

// Destroy the mute for specific user, return the user unmuted as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-mutes-users-destroy
func (s *MuteService) Destroy(params *MuteDestroyParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}
