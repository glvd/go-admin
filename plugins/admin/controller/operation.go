package controller

import (
	"encoding/json"
	"github.com/glvd/go-admin/modules/db"
	"github.com/glvd/go-admin/plugins/admin/models"

	"github.com/glvd/go-admin/context"
)

// RecordOperationLog record all operation logs, store into database.
func RecordOperationLog(ctx *context.Context) {
	if user, ok := ctx.UserValue["user"].(models.UserModel); ok {
		var input []byte
		form := ctx.Request.MultipartForm
		if form != nil {
			input, _ = json.Marshal((*form).Value)
		}

		models.OperationLog().SetConn(db.GetConnection(services)).New(user.Id, ctx.Path(), ctx.Method(), ctx.LocalIP(), string(input))
	}
}
