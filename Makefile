postgres-run:
	docker run --name=cargo-delivery-db -e POSTGRES_PASSWORD=$(password) -p 5432:5432 -d --rm postgres

migrate-create:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' down

run:
	go run cmd/main.go

osmr-create:
	docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-extract -p /opt/car.lua /data/russia-latest.osm.pbf
	docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-partition /data/volga-fed-district-latest.osrm
	docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-customize /data/volga-fed-district-latest.osrm

osmr-start:
	docker run -t -i --name OSRM_Backend_Car --restart=always -p 5000:5000 -v "${PWD}:/data" osrm/osrm-backend osrm-routed --algorithm mld /data/volga-fed-district-latest.osrm