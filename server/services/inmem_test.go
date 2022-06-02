package services

import (
	"testing"

	"github.com/galuhpradipta/go-inmem-db/server/store"
)

func Test_inMemHandler_set(t *testing.T) {
	mapStore := store.NewStore()

	type fields struct {
		store store.StoreHandler
	}
	type args struct {
		cmds []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "set #1",
			fields: fields{
				store: mapStore,
			},
			args: args{
				cmds: []string{"SET", "test_key_#1", "test_value_#1"},
			},
			want:    "set command success",
			wantErr: false,
		},
		{
			name: "set #2",
			fields: fields{
				store: mapStore,
			},
			args: args{
				cmds: []string{"SET", "test_key_#2", "test_value_#2"},
			},
			want:    "set command success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemHandler{
				store: tt.fields.store,
			}
			got, err := i.set(tt.args.cmds)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemHandler.set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inMemHandler.set() = %v, want %v", got, tt.want)
			}
		})
	}
}
