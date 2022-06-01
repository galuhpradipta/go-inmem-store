package store

import "testing"

type mock struct {
	Map map[string]string
}

func Test_dbHandler_Set(t *testing.T) {
	store := mock{
		Map: make(map[string]string),
	}

	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name      string
		mockStore mock
		args      args
	}{
		{
			name:      "Set #1",
			mockStore: store,
			args: args{
				key:   "test_key_#1",
				value: "test_value_#1",
			},
		},
		{
			name:      "Set #2",
			mockStore: store,
			args: args{
				key:   "test_key_#2",
				value: "test_value_#2",
			},
		},
	}
	for _, tt := range tests {
		db := &dbHandler{
			Map: tt.mockStore.Map,
		}

		t.Run(tt.name, func(t *testing.T) {
			db.Set(tt.args.key, tt.args.value)
			if tt.mockStore.Map[tt.args.key] != tt.args.value {
				t.Errorf("dbHandler.Get() = %v, want %v", db.Get(tt.args.key), "")
			}
		})
	}
}

func Test_dbHandler_Get(t *testing.T) {
	store := mock{
		Map: map[string]string{
			"test_key_#1": "test_value_#1",
			"test_key_#2": "test_value_#2",
		},
	}

	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name      string
		mockStore mock
		args      args
	}{
		{
			name:      "Get #1",
			mockStore: store,
			args: args{
				key:   "test_key_#1",
				value: "test_value_#1",
			},
		},
		{
			name:      "Get #2",
			mockStore: store,
			args: args{
				key:   "test_key_#2",
				value: "test_value_#2",
			},
		},
		{
			name:      "Get #3",
			mockStore: store,
			args: args{
				key:   "test_key_#3",
				value: "",
			},
		},
	}
	for _, tt := range tests {
		db := &dbHandler{
			Map: tt.mockStore.Map,
		}

		t.Run(tt.name, func(t *testing.T) {
			db.Get(tt.args.key)
			if tt.mockStore.Map[tt.args.key] != tt.args.value {
				t.Errorf("expect = %v, got %v", tt.args.value, db.Get(tt.args.key))
			}
		})
	}
}

func Test_dbHandler_Delete(t *testing.T) {
	store := mock{
		Map: map[string]string{
			"test_key_#1": "test_value_#1",
			"test_key_#2": "test_value_#2",
		},
	}

	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name      string
		mockStore mock
		args      args
	}{
		{
			name:      "Delete #1",
			mockStore: store,
			args: args{
				key:   "test_key_#1",
				value: "",
			},
		},
	}
	for _, tt := range tests {
		db := &dbHandler{
			Map: tt.mockStore.Map,
		}

		t.Run(tt.name, func(t *testing.T) {
			db.Delete(tt.args.key)
			if tt.mockStore.Map[tt.args.key] != tt.args.value {
				t.Errorf("expect = %v, got %v", tt.args.value, tt.mockStore.Map[tt.args.key])
			}
		})
	}
}
