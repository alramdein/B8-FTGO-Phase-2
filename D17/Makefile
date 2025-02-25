

mockgen:
	mockgen -destination=./mocks/mock_user_repository.go -package=mocks hacktiv/repository IUserRepository \
	&& mockgen -destination=./mocks/mock_user_usecase.go -package=mocks hacktiv/usecase IUserUsecase 

test:
	go test -cover -v ./...
