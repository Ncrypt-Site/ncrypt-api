package handlers

import (
	"github.com/labstack/echo"
	"ncrypt-api/config"
	"ncrypt-api/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDI_GetIndexV1(t *testing.T) {
	storage.Storage["shadow"] = storageShadowInterface{}
	c := config.BuildConfig()
	c.StorageDriver = "shadow"

	di, err := BuildDI(c)
	if err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	err = di.GetIndexV1(context)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fail()
	}
}
