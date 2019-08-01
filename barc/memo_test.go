package barc

import "testing"

func TestDecrypt(t *testing.T) {
	type args struct {
		msg     string
		fromPub string
		toPub   string
		wif     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			want: "hello boy",
			args: args{
				msg:     "e8e88ad5de168283c29a71635fe1d065",
				fromPub: "BAR6fpcoYK72BxsYYRwcBEPGuVoGhy2Yki2YCNfnCZCYxL5xp56Hh",
				toPub:   "BAR6icdz8dWibXRz8PcDn9RMupFkPbwHQ4toHxP8UmLm2hDtMHUKr",
				wif:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.msg, tt.args.fromPub, tt.args.toPub, tt.args.wif)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
