package repository

import "rank-master-back/internal/model/entity"

type IUser interface {
	// FindLockWithRankMasterAccountExist SELECT EXISTS(SELECT * FROM @@table WHERE rank_master_account = @rankMasterAccount FOR UPDATE)
	FindLockWithRankMasterAccountExist(rankMasterAccount string) (int64, error)
	// UpdateBatchUserByID
	// UPDATE @@table SET
	// name = CASE id
	// {{for _, user1 := range userList}}
	// WHEN @user1.ID THEN @user1.Name
	// {{end}}
	// END,
	// mobile = CASE id
	// {{for _, user2 := range userList}}
	// WHEN @user2.ID THEN @user2.Mobile
	// {{end}}
	// END
	// WHERE id IN @idList
	UpdateBatchUserByID(idList []string, userList []*entity.User) (int64, error)
}
