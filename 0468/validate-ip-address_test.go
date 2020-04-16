package _468

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{"172.16.254.1"}, want: T_IPv4},
		{name: "test1", args: args{"172.16.255.256"}, want: T_INVALID},
		{name: "test2", args: args{"2001:0db8:85a3:0:0:8A2E:0370:7334"}, want: T_IPv6},
		{name: "test2", args: args{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, want: T_IPv6},
		{name: "test2", args: args{"2001:db8:85a3:0:0:8A2E:0370:7334"}, want: T_IPv6},
		{name: "test2", args: args{"2001:0db8:85a3::8A2E:0370:7334"}, want: T_INVALID},
		{name: "test2", args: args{"02001:0db8:85a3:0000:0000:8a2e:0370:7334"}, want: T_INVALID},
		{name: "test3", args: args{"256.256.256.256"}, want: T_INVALID},
		{name: "test3", args: args{""}, want: T_INVALID},
		{name: "test3", args: args{"23:32.23.4"}, want: T_INVALID},
		{name: "test_1", args: args{"192.0.0.1"}, want: T_IPv4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
