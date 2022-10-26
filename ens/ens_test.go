package ens

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func TestENSStore_Resolve(t *testing.T) {
	type fields struct {
		backend      bind.ContractBackend
		resolverAddr string
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    common.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ensstore := &ENSStore{
					backend:      tt.fields.backend,
					resolverAddr: tt.fields.resolverAddr,
				}
				got, err := ensstore.Resolve(tt.args.input)
				if (err != nil) != tt.wantErr {
					t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Resolve() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
