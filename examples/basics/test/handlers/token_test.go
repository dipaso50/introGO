package handlers

import (
	"introGO/examples/basics/test/cache"
	"introGO/examples/basics/test/inmemory"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_create_token(t *testing.T) {

	tokCacheService := cache.NewTokenCacheService(inmemory.NewInmemoryCache())

	req, err := http.NewRequest("GET", "/token", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.Handler(CreateToken(tokCacheService))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	/**expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}**/

}
