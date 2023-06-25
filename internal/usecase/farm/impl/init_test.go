package impl

import (
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbPond "aquafarm/internal/repository/db/pond"
	uFarm "aquafarm/internal/usecase/farm"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		rDbFarm rDbFarm.Repository
		rDbPond rDbPond.Repository
	}
	tests := []struct {
		name string
		args args
		want uFarm.Usecase
	}{
		{
			name: "success",
			args: args{
				rDbFarm: &rDbFarm.MockRepository{},
				rDbPond: &rDbPond.MockRepository{},
			},
			want: &usecase{
				rDbFarm: &rDbFarm.MockRepository{},
				rDbPond: &rDbPond.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rDbFarm, tt.args.rDbPond); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
