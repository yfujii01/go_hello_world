package src

import (
	"testing"
	"os"
)

func TestFileWrite(t *testing.T) {
	type args struct {
		filename string
		mes      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "正常1", args: args{filename: "testfile1.txt", mes: "test1text"}, wantErr: false},
		{name: "正常2", args: args{filename: "testfile2.txt", mes: "test2text"}, wantErr: false},
		{name: "message empty", args: args{filename: "testfile.txt", mes: ""}, wantErr: false},
		{name: "filename empty", args: args{filename: "", mes: "testmessage"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FileWrite(tt.args.filename, tt.args.mes); (err != nil) != tt.wantErr {
				t.Errorf("FileWrite() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				//テストファイル削除
				if Exists(tt.args.filename) {
					os.Remove(tt.args.filename)
				}
			}
		})
	}
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
