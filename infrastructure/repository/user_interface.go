package repository

type IUser interface {
	// FindLockWithRankMasterAccountExist SELECT EXISTS(SELECT * FROM @@table WHERE rank_master_account = @rankMasterAccount FOR UPDATE)
	FindLockWithRankMasterAccountExist(rankMasterAccount string) (int64, error)
}
