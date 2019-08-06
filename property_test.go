package property

import "testing"

func TestCreate(t *testing.T) {
	type args struct {
		key  string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		// TODO: Add test cases.
		{
			name: "recoder 1",
			args: args{
				key:  "aaaaa",
				data: "{'sss':'dsdsd'}",
			},
			wantErr: false,
		},
		{
			name: "recoder 2",
			args: args{
				key:  "bbbbb",
				data: "{'kkkk':'dsdsd'}",
			},
			wantErr: false,
		},
		{
			name: "recoder 3",
			args: args{
				key:  "ccccc",
				data: "{'llll':'dsdsd'}",
			},
			wantErr: false,
		},
	}
	s := new(Storage)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(s, tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	showCurrentInfo(t)
}

func TestUpdate(t *testing.T) {
	type args struct {
		key  string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "change recoder 2",
			args: args{
				key:  "bbbbb",
				data: "{'kkkk':'dsdsda'}",
			},
			wantErr: false,
		},
	}
	s := new(Storage)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(s, tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	showCurrentInfo(t)
}

func TestRead(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Read key 2",
			args: args{
				key: "bbbbb",
			},
			want:    "{'kkkk':'dsdsda'}",
			wantErr: false,
		},
	}
	s := new(Storage)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(s, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
	showCurrentInfo(t)
}

func TestDelete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Delete key 2",
			args: args{
				key: "bbbbb",
			},
			want:    "{'kkkk':'dsdsda'}",
			wantErr: false,
		},
	}
	s := new(Storage)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(s, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
	showCurrentInfo(t)
}

func showCurrentInfo(t *testing.T) {
	keys := []string{"aaaaa", "bbbbb", "ccccc"}
	s := new(Storage)
	for i := range keys {
		v, e := Read(s, keys[i])
		if e == nil {
			t.Logf("Key '%s' Current data: %s", keys[i], v)
		}
	}
}
