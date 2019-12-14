package banana

import "errors"

var (
	RepoNotUpdated = errors.New("Cập nhật thông tin Repo thất bại")
	RepoNotFound   = errors.New("Repo không tồn tại")
	RepoConflict   = errors.New("Repo đã tồn tại")
	RepoInsertFail = errors.New("Thêm Repo thất bại")
	ErrorSql       = errors.New("Lỗi SQL")
)
