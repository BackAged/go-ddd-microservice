package config_test

import (
	"os"
	"testing"

	"github.com/BackAged/go-ddd-microservice/config"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		input   config.Application
		want    bool
		wantErr bool
	}{
		{
			input: config.Application{
				Port: 300,
				PostGreSQL: config.PostGreSQL{
					DBURL:     "root:root",
					DBTimeOut: 10,
				},
				RabbitMQ: config.RabbitMQ{
					RabbitURL: "rot:rot",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			input: config.Application{
				Port: 0,
				PostGreSQL: config.PostGreSQL{
					DBURL:     "root:root",
					DBTimeOut: 0,
				},
				RabbitMQ: config.RabbitMQ{
					RabbitURL: "",
				},
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run("Application IsValid", func(t *testing.T) {
			got, gotErr := tc.input.IsValid()

			if got != tc.want {
				t.Errorf("want '%t', got '%t'", tc.want, got)
			}

			if tc.wantErr == false && gotErr != nil {
				t.Errorf("want no error but got error")
			}

			if tc.wantErr == true && gotErr == nil {
				t.Errorf("want error but got no error")
			}

		})
	}
}

func TestGetApp(t *testing.T) {
	t.Run("GetApp", func(t *testing.T) {
		os.Setenv("DB_URL", "postgresql://root:root@localhost:5432/order?sslmode=disable")
		os.Setenv("DB_TIME_OUT", "100")
		os.Setenv("PORT", "3003")
		os.Setenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")

		_, gotErr := config.GetApp()
		if gotErr != nil {
			t.Errorf("want no error but error")
		}

		os.Clearenv()
	})
}
