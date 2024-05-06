package gempty

import (
	"reflect"
	"testing"
)

type foo struct {
	s string
	i int
	t bool
}

func TestClone(t *testing.T) {
	type args struct {
		s any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    any
	}{
		{"strig type", args{"1"}, false, ""},
		{"*strig type", args{pVar("a")}, false, pVar("")},
		{"int type", args{1}, false, 0},
		{"*int type", args{pVar(32)}, false, pVar(0)},
		{"bool type", args{true}, false, false},
		{"[]int type", args{[]int{0, 1}}, false, []int{}},
		{"[]string type", args{[]string{"a", "b"}}, false, []string{}},
		{"[2]int type", args{[2]int{0, 1}}, false, [2]int{0, 0}},
		{"[2]string type", args{[2]string{"a", "b"}}, false, [2]string{"", ""}},
		{"map[string]int type", args{map[string]int{"a": 1}}, false, map[string]int{}},
		{"struct type", args{foo{"a", 1, true}}, false, foo{"", 0, false}},
		{"*struct type", args{&foo{"a", 1, true}}, false, &foo{"", 0, false}},
		{"[]struct type", args{[]foo{{"a", 1, true}}}, false, []foo{}},
		{"chan type", args{make(chan int)}, true, nil},
		{"func type", args{func() {}}, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Clone(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Clone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				t.Logf("%T/%+v) --> (%T/%+v)\n", tt.args.s, tt.args.s, got, got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Clone got = %T/%+v, want: %T/%+v", got, got, tt.want, tt.want)
				}
			}
		})
	}
}

func pVar[T comparable](s T) *T {
	return &s
}
