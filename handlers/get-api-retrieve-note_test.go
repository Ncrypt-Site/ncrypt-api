package handlers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"ncrypt-api/config"
	"ncrypt-api/models"
	"ncrypt-api/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDI_GetSecureNoteV1(t *testing.T) {
	storage.Storage["shadow"] = storageShadowInterface{}
	c := config.BuildConfig()
	c.StorageDriver = "shadow"

	di, err := BuildDI(c)
	if err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//build context
	context := e.NewContext(req, rec)
	context.SetPath("api/v1/note/:id")
	context.SetParamNames("id")
	context.SetParamValues("wrong-uuid")

	err = di.GetSecureNoteV1(context)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fail()
	}

	responseModel := models.Response{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseModel)
	if err != nil {
		t.Fatal(err)
	}

	if responseModel.Message != "input failure" {
		t.Fail()
	}
}
