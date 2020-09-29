// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

type RequestIdHandler struct {
}

func (h *RequestIdHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.RequestIdKey)
	if err != nil {
		log.Printf("Problem reading RequestId from DB: %v\n", err)
		return 500, `{"error": "Could not read RequestId from DB"}`
	}
	return 200, fmt.Sprintf(`{ "RequestId": "%s" }`, string(v))
}
