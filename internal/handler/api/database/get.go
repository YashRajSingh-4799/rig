package database

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/rigdev/rig-go-api/api/v1/database"
	"github.com/rigdev/rig/pkg/uuid"
)

func (h *Handler) Get(ctx context.Context, req *connect.Request[database.GetRequest]) (*connect.Response[database.GetResponse], error) {
	dbId, err := uuid.Parse(req.Msg.GetDatabaseId())
	if err != nil {
		return nil, err
	}

	db, err := h.ds.Get(ctx, dbId)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&database.GetResponse{
		Database: db,
	}), nil
}
