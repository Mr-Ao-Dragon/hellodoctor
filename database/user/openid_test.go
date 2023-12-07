package database

import "testing"

func TestUpdateOpenId(t *testing.T) {
	type args struct {
		openId string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateOpenId(tt.args.openId); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOpenId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOpenId(t *testing.T) {
	type args struct {
		openId string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateOpenId(tt.args.openId); (err != nil) != tt.wantErr {
				t.Errorf("CreateOpenId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
