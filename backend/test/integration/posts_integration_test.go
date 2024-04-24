package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/doktorupnos/crow/backend/internal/post"
)

func TestPostsWorkflowIntegration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	token := createTestUser(t, "post_creator", "password", http.StatusCreated)
	defer deleteTestUser(t, "post_creator", "password")

	createTestPost(t, token, "1")
	createTestPost(t, token, "2")
	createTestPost(t, token, "3")
	assertEqual(t, len(getPosts(t, token, "page=0")), 3)

	createTestPost(t, token, "4")
	createTestPost(t, token, "5")
	posts := getPosts(t, token, "page=1")
	assertEqual(t, len(posts), 2)
}

func createTestPost(t testing.TB, token *http.Cookie, body string) {
	createPostReq, err := http.NewRequest(http.MethodPost, postsEndpoint, strings.NewReader(fmt.Sprintf(`
    {
    "body": %q
    }
    `, body)))
	if err != nil {
		t.Fatal(err)
	}
	createPostReq.AddCookie(token)
	resp, err := client.Do(createPostReq)
	if err != nil {
		t.Fatal(err)
	}

	assertStatusCode(t, resp.StatusCode, http.StatusCreated)
}

func getPosts(t testing.TB, token *http.Cookie, queryParameters ...string) []post.FeedPost {
	var queryString string
	if len(queryParameters) > 0 {
		queryString = "?" + strings.Join(queryParameters, "&")
	}
	url := postsEndpoint + queryString
	getPostsReq, err := http.NewRequest(http.MethodGet, url, noBody)
	if err != nil {
		t.Fatal(err)
	}
	getPostsReq.AddCookie(token)
	resp, err := client.Do(getPostsReq)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	assertStatusCode(t, resp.StatusCode, http.StatusOK)

	var posts []post.FeedPost
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		t.Fatal(err)
	}
	return posts
}
