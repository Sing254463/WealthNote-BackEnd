package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database Configuration
	DatabaseType     int  `env:"DATABASE_TYPE"`
	EnableMongoDB    bool `env:"ENABLE_MONGODB"`
	EnableMySQL      bool `env:"ENABLE_MYSQL"`
	EnableSQLServer  bool `env:"ENABLE_SQLSERVER"`
	EnablePostgreSQL bool `env:"ENABLE_POSTGRESQL"`

	// Database URLs
	DatabaseURL  string `env:"DATABASE_URL"`
	MongoDBURL   string `env:"MONGODB_URL"`
	SQLServerURL string `env:"SQLSERVER_URL"`

	// PostgreSQL Configuration
	PostgreSQLHost     string `env:"POSTGRESQL_HOST"`
	PostgreSQLPort     string `env:"POSTGRESQL_PORT"`
	PostgreSQLUser     string `env:"POSTGRESQL_USER"`
	PostgreSQLPassword string `env:"POSTGRESQL_PASSWORD"`
	PostgreSQLDatabase string `env:"POSTGRESQL_DATABASE"`
	PostgreSQLSSLMode  string `env:"POSTGRESQL_SSLMODE"`

	// MySQL Configuration
	MySQLHost     string `env:"MYSQL_HOST"`
	MySQLPort     string `env:"MYSQL_PORT"`
	MySQLUser     string `env:"MYSQL_USER"`
	MySQLPassword string `env:"MYSQL_PASSWORD"`
	MySQLDatabase string `env:"MYSQL_DATABASE"`

	// SQL Server Configuration
	SQLServerHost     string `env:"SQLSERVER_HOST"`
	SQLServerUser     string `env:"SQLSERVER_USER"`
	SQLServerPassword string `env:"SQLSERVER_PASSWORD"`
	SQLServerDB       string `env:"SQLSERVER_DB"`

	// MongoDB Configuration
	MongoDBURI string `env:"MONGODB_URI"`

	// JWT Configuration
	JWTSecret          string `env:"JWT_SECRET"`
	JWTExpirationShort string `env:"JWT_EXPIRATION_SHORT"`
	JWTExpirationLong  string `env:"JWT_EXPIRATION_LONG"`

	// OAuth Configuration
	EnableOAuth bool `env:"ENABLE_OAUTH"`

	// Server Configuration
	Port string `env:"PORT"`
}

var AppConfig *Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbType, _ := strconv.Atoi(getEnv("DATABASE_TYPE", "1"))

	AppConfig = &Config{
		DatabaseType:    dbType,
		EnableMongoDB:   getEnvAsBool("ENABLE_MONGODB", false),
		EnableMySQL:     getEnvAsBool("ENABLE_MYSQL", false),
		EnableSQLServer: getEnvAsBool("ENABLE_SQLSERVER", true),

		DatabaseURL:  getEnv("DATABASE_URL", ""),
		MongoDBURL:   getEnv("MONGODB_URL", ""),
		SQLServerURL: getEnv("SQLSERVER_URL", ""),

		PostgreSQLHost:     getEnv("POSTGRESQL_HOST", "localhost"),
		PostgreSQLPort:     getEnv("POSTGRESQL_PORT", "5432"),
		PostgreSQLUser:     getEnv("POSTGRESQL_USER", "postgres"),
		PostgreSQLPassword: getEnv("POSTGRESQL_PASSWORD", "password"),
		PostgreSQLDatabase: getEnv("POSTGRESQL_DATABASE", "knowledgeDB"),
		PostgreSQLSSLMode:  getEnv("POSTGRESQL_SSLMODE", "disable"),

		MySQLHost:     getEnv("MYSQL_HOST", "localhost"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLUser:     getEnv("MYSQL_USER", "root"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", "password"),
		MySQLDatabase: getEnv("MYSQL_DATABASE", "testdb"),

		SQLServerHost:     getEnv("SQLSERVER_HOST", "localhost"),
		SQLServerUser:     getEnv("SQLSERVER_USER", "sa"),
		SQLServerPassword: getEnv("SQLSERVER_PASSWORD", "password"),
		SQLServerDB:       getEnv("SQLSERVER_DB", "testdb"),

		MongoDBURI: getEnv("MONGODB_URI", "mongodb://localhost:27017/testdb"),

		JWTSecret:          getEnv("JWT_SECRET", "your_secret_key"),
		JWTExpirationShort: getEnv("JWT_EXPIRATION_SHORT", "15m"),
		JWTExpirationLong:  getEnv("JWT_EXPIRATION_LONG", "15d"),

		EnableOAuth: getEnvAsBool("ENABLE_OAUTH", false),

		Port: getEnv("PORT", "3000"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true"
	}
	return defaultValue
}
