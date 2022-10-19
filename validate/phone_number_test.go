package validate

import "testing"

func TestPhoneNumber(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
            name: "should pass",
			args: args{
				number: "+998913084678",
			},
			wantErr: false,
		},
		{
            name: "should fail",
			args: args{
				number: "9993084678",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PhoneNumber(tt.args.number); (err != nil) != tt.wantErr {
				t.Errorf("PhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
