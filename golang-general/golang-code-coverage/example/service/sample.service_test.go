package service

import (
	"reflect"
	"testing"
)

func TestCube_Volume(t *testing.T) {
	tests := []struct {
		name       string
		cube       Cube
		wantVolume int
	}{
		{
			name:       "Cube of size 1",
			cube:       Cube{1},
			wantVolume: 1,
		},
		{
			name:       "Cube of size 2",
			cube:       Cube{2},
			wantVolume: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVolume := tt.cube.Volume(); gotVolume != tt.wantVolume {
				t.Errorf("Cube.Volume() = %v, want %v", gotVolume, tt.wantVolume)
			}
		})
	}
}

func TestCube_SurfaceArea(t *testing.T) {
	tests := []struct {
		name     string
		cube     Cube
		wantArea int
	}{
		{
			name:     "Cube of size 1",
			cube:     Cube{1},
			wantArea: 6,
		},
		{
			name:     "Cube of size 2",
			cube:     Cube{2},
			wantArea: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArea := tt.cube.SurfaceArea(); gotArea != tt.wantArea {
				t.Errorf("Cube.SurfaceArea() = %v, want %v", gotArea, tt.wantArea)
			}
		})
	}
}

func TestNewCube(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *Cube
	}{
		{
			name: "Cube of size 1",
			args: args{1},
			want: &Cube{1},
		},
		{
			name: "Cube of size 2",
			args: args{2},
			want: &Cube{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCube(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCube() = %v, want %v", got, tt.want)
			}
		})
	}
}
