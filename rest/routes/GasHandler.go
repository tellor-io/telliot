package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

//BalanceHandler handles balance requests
type GasHandler struct {
}

//Incoming implementation for  handler
func (h *GasHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.GasKey)
	if err != nil {
		log.Printf("Problem reading Gas from DB: %v\n", err)
		return 500, `{"error": "Could not read Gas from DB"}`
	}
	return 200, fmt.Sprintf(`{ "Gas": "%s" }`, string(v))
}
