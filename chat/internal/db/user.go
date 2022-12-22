package db

import (
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
)

func RegisterNewUser(request *constants.NewUserCreated) {
	_, err := execute(nil, InsertUserQuery(request.User_id.Int64(), request.Name))
	if err != nil {
		logger.Logger.Sugar().Error("error registering new user:", err)
	}
}
