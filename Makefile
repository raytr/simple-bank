build:
	docker build --tag simple-bank .

mocks:
	@echo "Generate mock repository"
	cd repository && mockery --all --case underscore --output repository_mock --outpkg repository_mock

serve-serve:
	cd hackathon/backend && docker-compose up


up-migration:
	migrate -path migrations -database "postgresql://postgres:postgres@127.0.0.1:5432/account_authentication_db?sslmode=disable" -verbose up

create-network:
	docker network create account-authentication-network

repository-mocks:
	@echo "Generate mock repository"
	cd repository && mockery --all --case underscore --output repository_mock --outpkg repository_mock

service-mocks:
	@echo "Generate mock endpoints"
	cd services && mockery --all --case underscore --output services_mock --outpkg services_mock

