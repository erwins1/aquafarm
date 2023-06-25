package impl

import (
	"aquafarm/common/dbw"
	rDbPond "aquafarm/internal/repository/db/pond"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		db dbw.Repository
	}
	tests := []struct {
		name string
		args args
		want rDbPond.Repository
	}{
		{
			name: "success",
			args: args{
				&dbw.MockRepository{},
			},
			want: &repository{
				&dbw.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
