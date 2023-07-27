package shop

import (
	"reflect"
	"testing"
)

func TestNewShop(t *testing.T) {
	type args struct {
		ownerID     string
		name        string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    *Shop
		wantErr bool
	}{
		// todo impl
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewShop(tt.args.ownerID, tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewShop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newShop(t *testing.T) {
	type args struct {
		id          string
		ownerID     string
		name        string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    *Shop
		wantErr bool
	}{
		// todo impl
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newShop(tt.args.id, tt.args.ownerID, tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("newShop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newShop() = %v, want %v", got, tt.want)
			}
		})
	}
}
