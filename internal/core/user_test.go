package core

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/db/mock"

// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/model"
// 	"github.com/sayanibhattacharjee/sequoia-backend-assignment/internal/repository"
// )

// func TestUserCore_GetByID(t *testing.T) {
// 	user := model.User{
// 		ID:   1,
// 		Name: "Vishnu",
// 	}
// 	mockRepo := mock.UserMockRepository{
// 		DB: map[string]*model.User{
// 			"1": &user,
// 		},
// 	}

// 	type fields struct {
// 		ur repository.UserRepository
// 	}
// 	type args struct {
// 		id string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *model.User
// 		wantErr bool
// 	}{
// 		{
// 			"must return user model",
// 			fields{
// 				&mockRepo,
// 			},
// 			args{"1"},
// 			&user,
// 			false,
// 		},
// 		{
// 			"must return nil if user does not exist",
// 			fields{
// 				&mockRepo,
// 			},
// 			args{"2"},
// 			nil,
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &UserCore{
// 				ur: tt.fields.ur,
// 			}
// 			got, err := c.GetByID(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserCore.GetByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UserCore.GetByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
