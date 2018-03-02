package arbitrary

import (
	"reflect"
	"testing"
)

func anonimize(data interface{}) interface{} {
	var anon interface{}
	err := Hydrate(data, &anon)
	if err != nil {
		panic(err)
	}
	return anon
}

func TestJoiner_flatten(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		j       Joiner
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "bare string",
			j:       DefaultJoiner,
			args:    args{data: "foobar"},
			want:    map[string]interface{}{"": "foobar"},
			wantErr: false,
		},
		{
			name: "string array",
			j:    DefaultJoiner,
			args: args{data: anonimize([]string{"foo", "bar"})},
			want: map[string]interface{}{
				"0": "foo",
				"1": "bar",
			},
			wantErr: false,
		},
		{
			name: "string map",
			j:    DefaultJoiner,
			args: args{data: anonimize(map[string]string{
				"foo": "bar",
				"baz": "foobar",
			})},
			want: map[string]interface{}{
				"foo": "bar",
				"baz": "foobar",
			},
			wantErr: false,
		},
		{
			name: "nested string map",
			j:    DefaultJoiner,
			args: args{data: anonimize(
				map[string]interface{}{
					"foo":    "bar",
					"answer": 42,
					"baz": map[string]string{
						"jinx": "jax",
					},
				})},
			want: map[string]interface{}{
				"foo":      "bar",
				"answer":   42.0,
				"baz.jinx": "jax",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.flatten(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Joiner.flatten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Joiner.flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "bare string",
			args:    args{data: "foobar"},
			want:    map[string]interface{}{"": "foobar"},
			wantErr: false,
		},
		{
			name: "string array",
			args: args{data: anonimize([]string{"foo", "bar"})},
			want: map[string]interface{}{
				"0": "foo",
				"1": "bar",
			},
			wantErr: false,
		},
		{
			name: "string map",
			args: args{data: anonimize(map[string]string{
				"foo": "bar",
				"baz": "foobar",
			})},
			want: map[string]interface{}{
				"foo": "bar",
				"baz": "foobar",
			},
			wantErr: false,
		},
		{
			name: "nested string map",
			args: args{data: anonimize(
				map[string]interface{}{
					"foo":    "bar",
					"answer": 42,
					"baz": map[string]string{
						"jinx": "jax",
					},
				})},
			want: map[string]interface{}{
				"foo":      "bar",
				"answer":   42.0,
				"baz.jinx": "jax",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Flatten(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Flatten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
