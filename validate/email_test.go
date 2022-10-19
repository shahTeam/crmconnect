package validate

import "testing"

func TestEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{email: "shahzodasturchi@gmail.com"},
			wantErr: false,
		},
		{
            name: "should fail",
			args: args{email: "shahzodasturchigmail.com"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Email(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("Email() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
