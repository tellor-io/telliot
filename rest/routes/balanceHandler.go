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
type BalanceHandler struct {
}

//Incoming implementation for  handler
func (h *BalanceHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.BalanceKey)
	if err != nil {
		log.Printf("Problem reading balance from DB: %v\n", err)
		return 500, `{"error": "Could not read balance from DB"}`
	}
	return 200, fmt.Sprintf(`{ "balance": "%s" }`, string(v))
}
