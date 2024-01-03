package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/eliasmanj/budgets-api/db/mock"
	db "github.com/eliasmanj/budgets-api/db/sqlc"
	"github.com/eliasmanj/budgets-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockQuerier(ctrl)
	dbuser, pwd := randomUser(t)

	arg := db.CreateUserParams{
		Username: dbuser.Username,
		Email:    dbuser.Email,
	}
	store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(arg, pwd)).
		Times(1).Return(dbuser, nil)

	body := gin.H{
		"username": dbuser.Username,
		"password": pwd,
		"email":    dbuser.Email,
	}

	data, err := json.Marshal(body)
	require.NoError(t, err)

	url := "/users"
	server := newTestServer(t, store)
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	var response userReponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	require.Equal(t, dbuser.Username, response.Username)
	require.Equal(t, dbuser.Email, response.Email)

}

func TestGetUser(t *testing.T) {

}

func randomUser(t *testing.T) (user db.User, password string) {
	password = utils.RandomString(6)
	hashedPassword, err := utils.HashedPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       utils.RandomUser(),
		HashedPassword: hashedPassword,
		Email:          utils.RandomEmail(),
	}
	return
}

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := utils.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}
