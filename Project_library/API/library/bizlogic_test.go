package library

import (
	"Project_library/model"
	"database/sql"
	"reflect"
	"testing"
)

func TestGetMembers_DB(t *testing.T) {
	type args struct {
		db      *sql.DB
		members []model.Member
	}
	tests := []struct {
		name string
		args args
		want []model.Member
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMembers_DB(tt.args.db, tt.args.members); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMembers_DB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelMembers_DB(t *testing.T) {
	type args struct {
		db *sql.DB
		id string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelMembers_DB(tt.args.db, tt.args.id)
		})
	}
}

func TestGetMember_DB(t *testing.T) {
	type args struct {
		db      *sql.DB
		members []model.Member
		id      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetMember_DB(tt.args.db, tt.args.members, tt.args.id)
		})
	}
}

func TestAddMember_DB(t *testing.T) {
	type args struct {
		db     *sql.DB
		member model.Member
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddMember_DB(tt.args.db, tt.args.member)
		})
	}
}

func TestUpdateMember_DB(t *testing.T) {
	type args struct {
		db     *sql.DB
		member model.Member
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateMember_DB(tt.args.db, tt.args.member)
		})
	}
}
