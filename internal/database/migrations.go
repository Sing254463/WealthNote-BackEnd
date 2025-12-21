package database

import (
	"database/sql"
	"log"
)

func RunMigrations() {
	if postgresDB != nil {
		migratePostgreSQL(postgresDB)
	}
}

func migratePostgreSQL(db *sql.DB) {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS public.users
    (
        id_user SERIAL PRIMARY KEY,
        usercode VARCHAR(255),
        email VARCHAR(255),
        fnamet VARCHAR(255),
        lnamet VARCHAR(255),
        fnamee VARCHAR(255),
        lnamee VARCHAR(255),
        password TEXT NOT NULL,
        provider VARCHAR(100),
        provider_id TEXT,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
    );
    
    CREATE INDEX IF NOT EXISTS idx_users_email ON public.users(email);
    CREATE INDEX IF NOT EXISTS idx_users_usercode ON public.users(usercode);
    
    COMMENT ON TABLE public.users IS 'เป็น Tables สำหรับเก็บข้อมูลของ ผู้ใช้';
    `
	// CREATE INDEX IF NOT EXISTS idx_users_email ON public.users(email);
	// CREATE INDEX IF NOT EXISTS idx_users_usercode ON public.users(usercode);
	// เป็นการเพิ่มดัชนี (Index) บนคอลัมน์ email และ usercode ในตาราง users เพื่อเพิ่มประสิทธิภาพในการค้นหาข้อมูลตามคอลัมน์เหล่านี้
	_, err := db.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
	log.Println("✓ Users table migrated successfully")
}
