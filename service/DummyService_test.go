package service

import (
	"reflect"
	"testing"

	"github.com/DivTwid/golang-boiler/config"
	"github.com/DivTwid/golang-boiler/dto"
	"github.com/DivTwid/golang-boiler/model"
	"gorm.io/gorm"
)

func Test_dummyService_GetVal(t *testing.T) {
	tests := []struct {
		name string
		ds   dummyService
		want string
	}{
		{
			name: "test_service",
			ds:   dummyService{},
			want: "pong",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dummyService{}
			if got := ds.GetVal(); got != tt.want {
				t.Errorf("dummyService.GetVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dummyService_AddVal(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user dto.UserDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.User
	}{
		{
			name: "test add user",
			fields: fields{
				db: config.PqDB,
			},
			args: args{
				user: dto.UserDto{
					Name:    "test",
					Email:   "test@test.com",
					PhoneNo: "135212341",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := dummyService{
				db: tt.fields.db,
			}
			if got := ds.AddVal(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dummyService.AddVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
