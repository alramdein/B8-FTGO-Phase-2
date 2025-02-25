package handler

import (
	"encoding/json"
	"hacktiv/mocks"
	"hacktiv/model"
	"hacktiv/usecase"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := model.User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@.com",
	}

	userJSON, err := json.Marshal(user)
	require.NoError(t, err)

	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder() // simulasi request
	c := e.NewContext(req, rec)

	mockUserRepo := mocks.NewMockIUserRepository(ctrl)
	userUsecase := usecase.NewUserUsecase(mockUserRepo, nil)
	h := NewUserHandler(userUsecase)

	// mockUserRepo.On("CreateUser", mock.Anything, user).Return(nil) // ini pake mock testfy
	mockUserRepo.EXPECT().CreateUser(gomock.Any(), user).Return(nil) // ini pake gomock & mockgen

	err = h.CreateUserHandler(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
