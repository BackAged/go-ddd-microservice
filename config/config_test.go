package config_test

import (
	"log"
	"testing"

	"github.com/BackAged/go-ddd-microservice/config"
	"github.com/joho/godotenv"
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
		err := godotenv.Load(".env.example")
		if err != nil {
			log.Fatal("Error loading .env.examplee file")
		}

		_, gotErr := config.GetApp()
		if gotErr != nil {
			t.Errorf("want no error but error")
		}
	})
}
