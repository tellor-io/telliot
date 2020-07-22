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
type TotalTipHandler struct {
}

//Incoming implementation for  handler
func (h *TotalTipHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.TotalTipKey)
	if err != nil {
		log.Printf("Problem reading TotalTip from DB: %v\n", err)
		return 500, `{"error": "Could not read TotalTip from DB"}`
	}
	return 200, fmt.Sprintf(`{ "TotalTip": "%s" }`, string(v))
}
