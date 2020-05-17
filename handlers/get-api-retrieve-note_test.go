package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"ncrypt-api/config"
	"ncrypt-api/models"
	"ncrypt-api/storage"
	redisStorage "ncrypt-api/storage/redis-storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

type storageShadowWithFailedExists struct {
	redisStorage.RedisStorage
}

func (s storageShadowWithFailedExists) BuildConfiguration(c models.Config) (models.StorageInterface, error) {
	return s, nil
}

func (s storageShadowWithFailedExists) Exists(id uuid.UUID) bool {
	return false
}

func TestDI_GetSecureNoteV1WithInvalidUUID(t *testing.T) {
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

func TestDI_GetSecureNoteV1WithStorageFailure(t *testing.T) {
	storage.Storage["shadow"] = storageShadowWithFailedExists{}
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

	// injecting validator
	v := validator.New()
	e.Validator = &customValidator{validator: v}

	//build context
	context := e.NewContext(req, rec)
	context.SetPath("api/v1/note/:id")
	context.SetParamNames("id")
	context.SetParamValues(uuid.New().String())

	err = di.GetSecureNoteV1(context)
	if err != nil {
		t.Fatal(err)
	}

	responseModel := models.Response{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseModel)
	if err != nil {
		t.Fatal(err)
	}

	if responseModel.Message != "unable to retrieve note" {
		t.Fail()
	}
}
