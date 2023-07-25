package config

import "testing"

func TestPostgresDB_InitializeDB(t *testing.T) {
	type fields struct {
		host     string
		user     string
		password string
		dbname   string
		port     string
		sslmode  string
		timezone string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "postgres_test",
			fields: fields{
				host:     "localhost",
				user:     "root",
				password: "root",
				dbname:   "postgres",
				port:     "5432",
				sslmode:  "disable",
				timezone: "Asia/Shanghai",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := PostgresDB{
				host:     tt.fields.host,
				user:     tt.fields.user,
				password: tt.fields.password,
				dbname:   tt.fields.dbname,
				port:     tt.fields.port,
				sslmode:  tt.fields.sslmode,
				timezone: tt.fields.timezone,
			}
			ps.InitializeDB()
		})
	}
}
