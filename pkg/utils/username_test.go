package utils

import "testing"

func TestIsValidUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "Dot in the start",
			args: args{username: ".asdasd"},
			want: false,
		},
		{
			name: "Dot in the end",
			args: args{username: "asdasd."},
			want: false,
		},
		{
			name: "Two underscores in a row",
			args: args{username: "mkil__ohk"},
			want: true,
		},
		{
			name: "Correct username with dot and underscore",
			args: args{username: "mkil._ohk"},
			want: true,
		},
		{
			name: "Two dots in a row",
			args: args{username: "mkil..ohk"},
			want: false,
		},
		{
			name: "Uppercase letter in username",
			args: args{username: "mkilOhk"},
			want: false,
		},
		{
			name: "Username with space",
			args: args{username: "mkil ohk"},
			want: false,
		}, {
			name: "Username with length 50",
			args: args{username: "mkilohkmkilohkmkilohkmkilohkmkilohk"},
			want: true,
		},
		{
			name: "Username with length greater than 50",
			args: args{username: "mkilohkmkilohkmkiloadasdasdasdasdhkmkilohkmkilohkmkilohk"},
			want: false,
		},
		{
			name: "Username with capital letter",
			args: args{username: "mkilKhk"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUsername(tt.args.username); got != tt.want {
				t.Errorf("IsValidUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
