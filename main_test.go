package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestRouter(t *testing.T) {
	r := newRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	mockServer := httptest.NewServer(r)

	// We make a GET request to the "hello" route we defined in the router
	resp, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}

	// We expect status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	// read the body into a bunch of bytes(b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := "Hello world!"

	// We want our response to match the one defined in our handler. If it does
	// happen to be "Hello world!", then it confirms, that the route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	mockServer := httptest.NewServer(r)

	// We make a POST request to the "hello" route that we didn't define
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We expect status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be %d, got %d",
			http.StatusMethodNotAllowed, resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// We want to hit the `GET /static/` route to get the index.html file
	resp, err := http.Get(mockServer.URL + "/static/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %d, got %d", http.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expected := "text/html; charset=utf-8"

	if contentType != expected {
		t.Errorf("Wrong content type, want %s, got %s", expected, contentType)
	}
}
