package gen

type IUser interface {
	// FindWithMobile SELECT EXISTS(SELECT * FROM @@table WHERE mobile = @mobile)
	FindWithMobile(mobile string) (int64, error)
}
