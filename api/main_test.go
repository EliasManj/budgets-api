package api

import (
	"os"
	"testing"

	db "github.com/eliasmanj/budgets-api/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Querier) *Server {
	server, err := NewServer(store)

	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
