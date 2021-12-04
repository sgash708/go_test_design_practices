package mock

import (
	"github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/model"
	"github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/repository"
)

// REF: https://moneyforward.com/engineers_blog/2021/03/08/go-test-mock/
// uRepoMock ユーザリポジトリをモックする構造体
type UserRepoMock struct {
	repository.User
	FakeCreate func(*model.User) (*model.User, error)
}

/**
 * モックの肝：
 * Fake*関数のフィールドをstructに追加する
 * → structのメソッドに求められるメソッドを実装
 * ※メソッドの定義がインターフェースと異なるとエラーを吐く
 */
// Create ユーザ作成モック
func (m *UserRepoMock) Create(user *model.User) (*model.User, error) {
	return m.FakeCreate(user)
}
