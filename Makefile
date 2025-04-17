# Menjalankan service utama
run:
	go run cmd/app/*.go

# Menjalankan migrasi database (up)
migrate:
	go run cmd/migration/*.go -type=run

# Menjalankan rollback migrasi (down)
rollback:
	go run cmd/migration/*.go -type=rollback

