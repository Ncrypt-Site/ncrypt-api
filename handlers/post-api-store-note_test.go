package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"ncrypt-api/config"
	"ncrypt-api/models"
	"ncrypt-api/storage"
	redisStorage "ncrypt-api/storage/redis-storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	postInvalidPayload = `{"message":"this is my branch","self_destruct": 2,"destruct_after_opening": true}`
	postValidPayload   = `{"message":"this is my branch","self_destruct": 0,"destruct_after_opening": true}`
)

type customValidator struct {
	validator *validator.Validate
}

func (v customValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

type storageShadowWithFailedStore struct {
	redisStorage.RedisStorage
}

func (s storageShadowWithFailedStore) BuildConfiguration(c models.Config) (models.StorageInterface, error) {
	return s, nil
}

func (s storageShadowWithFailedStore) Store(id uuid.UUID, data []byte, duration time.Duration) error {
	return errors.New("an error")
}

func (s storageShadowInterface) Store(id uuid.UUID, data []byte, duration time.Duration) error {
	return nil
}

func TestDI_PostStoreSecureNoteV1WithNoValidData(t *testing.T) {
	storage.Storage["shadow"] = storageShadowInterface{}
	c := config.BuildConfig()
	c.StorageDriver = "shadow"

	di, err := BuildDI(c)
	if err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	err = di.PostStoreSecureNoteV1(context)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fail()
	}
}

func TestDI_PostStoreSecureNoteV1WithInvalidPayload(t *testing.T) {
	storage.Storage["shadow"] = storageShadowInterface{}
	c := config.BuildConfig()
	c.StorageDriver = "shadow"

	di, err := BuildDI(c)
	if err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postInvalidPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// injecting validator
	v := validator.New()
	e.Validator = &customValidator{validator: v}

	//build context
	context := e.NewContext(req, rec)

	err = di.PostStoreSecureNoteV1(context)
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

	if responseModel.Message != "validation failed" ||
		responseModel.Error[0] != "validation failed for field: SelfDestruct. "+
			"reason: oneof. additional data: 0 1 3 6 12 24 48 72 168 720" {
		t.Fail()
	}
}

func TestDI_PostStoreSecureNoteV1WithValidPayloadAndStorageInterfaceFailure(t *testing.T) {
	storage.Storage["shadowWithFailedStore"] = storageShadowWithFailedStore{}
	c := config.BuildConfig()
	c.StorageDriver = "shadowWithFailedStore"

	di, err := BuildDI(c)
	if err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postValidPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// injecting validator
	v := validator.New()
	e.Validator = &customValidator{validator: v}

	//build context
	context := e.NewContext(req, rec)

	err = di.PostStoreSecureNoteV1(context)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusInternalServerError {
		t.Fail()
	}

	responseModel := models.Response{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseModel)
	if err != nil {
		t.Fatal(err)
	}

	if responseModel.Message != "internal error occurred" {
		t.Fail()
	}
}

func TestDI_PostStoreSecureNoteV1WithValidPayload(t *testing.T) {
	storage.Storage["shadow"] = storageShadowInterface{}
	c := config.BuildConfig()
	c.StorageDriver = "shadow"

	di, err := BuildDI(c)
	if err != nil {
		t.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postValidPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// injecting validator
	v := validator.New()
	e.Validator = &customValidator{validator: v}

	//build context
	context := e.NewContext(req, rec)

	err = di.PostStoreSecureNoteV1(context)
	if err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusCreated {
		t.Fail()
	}

	responseModel := models.Response{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseModel)
	if err != nil {
		t.Fatal(err)
	}

	if responseModel.Message != "Note stored." {
		t.Fail()
	}

	if responseModel.Data.(map[string]interface{})["id"] == "" {
		t.Fail()
	}
}
