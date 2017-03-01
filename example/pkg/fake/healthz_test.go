package fake

import "testing"

func TestHealthz(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{ "test #1", "I'm OK!", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Healthz()
			if (err != nil) != tt.wantErr {
				t.Errorf("Healthz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Healthz() = %v, want %v", got, tt.want)
			}
		})
	}
}
