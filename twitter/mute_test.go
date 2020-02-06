package twitter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMuteService_CreateService(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/mutes/users/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "golang"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"screen_name": "golang"}`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Mute.Create(&MuteCreateParams{ScreenName: "golang"})
	expected := User{ScreenName: "golang"}
	assert.Nil(t, err)
	assert.Equal(t, expected, users)
}

func TestMuteService_DestroyService(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/mutes/users/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "golang"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"screen_name": "golang"}`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Mute.Destroy(&MuteDestroyParams{ScreenName: "golang"})
	expected := User{ScreenName: "golang"}
	assert.Nil(t, err)
	assert.Equal(t, expected, users)
}
