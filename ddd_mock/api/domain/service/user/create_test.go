package user

import (
	"testing"

	"github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/model"
	"github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/service/user/mock"
)

// TestCreate_Success ユーザ作成_成功
func TestCreate_Success(t *testing.T) {
	cases := []struct {
		ID      int
		Name    string
		Address string
	}{
		{ID: 1, Name: "fakename", Address: "1020-20"},
		{ID: 2, Name: "fakenamae", Address: "1020-21"},
		{ID: 3, Name: "faketaro", Address: "999-99"},
		{ID: 4, Name: "fakeko", Address: "1029-777"},
	}
	userRepo := &mock.UserRepoMock{
		FakeCreate: func(user *model.User) (*model.User, error) {
			created := &model.User{ID: user.ID, Name: user.Name, Address: user.Address}
			return created, nil
		},
	}

	for _, c := range cases {
		t.Run("作成成功", func(t *testing.T) {
			userService := NewUser(userRepo)

			userInput := model.User{ID: c.ID, Name: c.Name, Address: c.Address}
			userData, err := userService.Create(&userInput)
			if err != nil {
				t.Fatal(err)
			}

			if userData.ID != c.ID {
				t.Errorf("User.Create()は、model.User.IDが間違っています。\n実際の入力：%d", userData.ID)
			}
			if userData.Name != userInput.Name {
				t.Errorf("User.Create()は、model.User.Nameが間違っています。\n実際の入力：%s", userData.Name)
			}
			if userData.Address != userInput.Address {
				t.Errorf("User.Create()は、model.User.Addressが間違っています。\n実際の入力：%s", userData.Address)
			}
		})
	}
}
