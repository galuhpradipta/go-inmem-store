package services

import (
	"testing"

	"github.com/galuhpradipta/go-inmem-db/server/store"
)

type fields struct {
	store store.StoreHandler
}

type args struct {
	cmds []string
}

type test struct {
	name    string
	fields  fields
	args    args
	want    string
	wantErr bool
}

func Test_inMemHandler_set(t *testing.T) {
	mapStore := store.NewStore()

	tests := []test{
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
		{
			name: "invalid command",
			fields: fields{
				store: mapStore,
			},
			args: args{
				cmds: []string{"SET", "test_key_#2"},
			},
			wantErr: true,
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

func Test_inMemHandler_dump(t *testing.T) {
	memStore := store.NewStore()
	memStore.Set("test_key_#1", "test_value_#1")
	memStore.Set("test_key_#2", "test_value_#2")

	tests := []test{
		{
			name: "dump #1",
			fields: fields{
				store: memStore,
			},
			args: args{
				cmds: []string{"DUMP", "test_key_#1"},
			},
			want: "test_value_#1",
		},
		{
			name: "dump #2",
			fields: fields{
				store: memStore,
			},
			args: args{
				cmds: []string{"DUMP", "test_key_#2"},
			},
			want: "test_value_#2",
		},

		{
			name: "invalid command",
			fields: fields{
				store: memStore,
			},
			args: args{
				cmds: []string{"DUMP", "test_key_#2", "test_value_#2"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemHandler{
				store: tt.fields.store,
			}
			got, err := i.dump(tt.args.cmds)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemHandler.dump() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inMemHandler.dump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemHandler_rename(t *testing.T) {
	memStore := store.NewStore()
	memStore.Set("test_key_#1", "test_value_#1")
	memStore.Set("test_key_#2", "test_value_#2")

	tests := []test{
		{
			name: "rename #1",
			fields: fields{
				store: memStore,
			},
			args: args{
				cmds: []string{"RENAME", "test_key_#1", "test_key_#2"},
			},
			want: "rename command success",
		},
		{
			name: "not found",
			fields: fields{
				store: memStore,
			},
			args: args{
				cmds: []string{"RENAME", "test_key_#3", "test_key_#4"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemHandler{
				store: tt.fields.store,
			}
			got, err := i.rename(tt.args.cmds)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemHandler.rename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inMemHandler.rename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemHandler_delete(t *testing.T) {
	mapStore := store.NewStore()
	mapStore.Set("test_key_#1", "test_value_#1")

	tests := []test{
		{
			name: "delete #1",
			fields: fields{
				store: mapStore,
			},
			args: args{
				cmds: []string{"DELETE", "test_key_#1"},
			},
			want: "delete command success",
		},
		{
			name: "invalid command",
			fields: fields{
				store: mapStore,
			},
			args: args{
				cmds: []string{"DELETE", "test_key_#1", "test_value_#1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemHandler{
				store: tt.fields.store,
			}
			got, err := i.delete(tt.args.cmds)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemHandler.delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inMemHandler.delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
