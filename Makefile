d:
	docker compose down
up:
	docker compose up -d
log:
	docker compose logs -f
ps:
	docker compose ps
mysql:
	docker exec -it HackU2023_Nagoya_DB mysql -u root --password=admin
go:
	docker exec -it HackU2023_Nagoya_GO ash
voicevox:
	docker exec -it HackU2023_Nagoya_Voicevox bash
