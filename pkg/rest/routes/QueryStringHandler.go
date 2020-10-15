// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tellor-io/TellorMiner/pkg/common"
	"github.com/tellor-io/TellorMiner/pkg/db"
)

type QueryStringHandler struct {
}

func (h *QueryStringHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.QueryStringKey)
	if err != nil {
		log.Printf("Problem reading QueryString from DB: %v\n", err)
		return 500, `{"error": "Could not read QueryString from DB"}`
	}
	return 200, fmt.Sprintf(`{ "QueryString": "%s" }`, string(v))
}
