package controller

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginExecutesCorrectTemplate(t *testing.T) {
	auth := new(authentication)
	expected := `login template`
	auth.loginTemplate, _ = template.New("").Parse(expected)
	r := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()

	auth.handleLogin(w, r)

	actual, _ := ioutil.ReadAll(w.Result().Body)

	if string(actual) != expected {
		t.Errorf("Failed  execute correct template")
	}
}
