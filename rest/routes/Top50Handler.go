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
type Top50Handler struct {
}

//Incoming implementation for  handler
func (h *Top50Handler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.Top50Key)
	if err != nil {
		log.Printf("Problem reading Top50 from DB: %v\n", err)
		return 500, `{"error": "Could not read Top50 from DB"}`
	}
	return 200, fmt.Sprintf(`{ "Top50": "%s" }`, string(v))
}
