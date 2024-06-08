package dao

type IUser interface {
	// FindLockWithRankMasterAccount SELECT EXISTS(SELECT * FROM @@table WHERE rank_master_account = @rankMasterAccount FOR UPDATE)
	FindLockWithRankMasterAccount(rankMasterAccount string) (int64, error)
}
