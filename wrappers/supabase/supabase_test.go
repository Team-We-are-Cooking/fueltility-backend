package fueltilitysupabase

import (
	"testing"

	"github.com/joho/godotenv"
)

func Test_SupabaseConnection(t *testing.T) {
    t.Run("Test connection to supabase postgreSQL database", func(t *testing.T) {
        if err := godotenv.Load("../../.env"); err != nil {
            t.Fatalf("Unable to load env file: %s", err.Error())
        }

        if _, err := CreateClient(); err != nil {
            t.Errorf("Unable to establish connection to database: %s", err.Error())
        }
    })

    t.Run("No url provided", func(t *testing.T) {
        t.Setenv("SUPABASE_URL", "")

        if _, err := CreateClient(); err == nil {
            t.Errorf("Expected error SUPABASE_URL is required")
        }
    })

    t.Run("No key provided", func(t *testing.T) {
        t.Setenv("SUPABASE_KEY", "")

        if _, err := CreateClient(); err == nil {
            t.Errorf("Expected error SUPABASE_KEY is required")
        }
    })
}