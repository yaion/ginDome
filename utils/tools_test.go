package utils

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestMd5str(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{str: "123456"},
			want: "e10adc3949ba59abbe56e057f20f883e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5str(tt.args.str); got != tt.want {
				t.Errorf("Md5str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateVerifyImg(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "test",
			args: args{height: 100, width: 200},
			want: map[string]string{
				"id":  "test",
				"img": "iVBORw0KGgoAAAANSUhEUgAAAIAAAACACAYAAADDPmHLAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAadEVYdFNvZnR3YX",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateVerifyImg(tt.args.height, tt.args.width)
			fmt.Println(got)
			time.Sleep(10 * time.Minute)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateVerifyImg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateVerifyImg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		id     string
		digits string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{id: "85vnif4DGmvdK4MdglPz", digits: "755121"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verify(tt.args.id, tt.args.digits); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
