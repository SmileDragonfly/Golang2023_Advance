package main

import (
	"context"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:123@123A@192.168.68.242:1433?database=gormdb"))
	if err != nil {
		panic(err)
	}
	return db
}

func TestQueryGorm_GetAllUsers(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []User
		wantErr bool
	}{
		{
			name: "TestQueryGorm_GetAllUsers",
			fields: fields{
				db: setupDB(),
			},
			args: args{
				ctx: context.TODO(),
			},
			want: []User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QueryGorm{
				db: tt.fields.db,
			}
			got, err := q.GetAllUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
