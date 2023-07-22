package mock

//go:generate mockgen -package=domain_test -destination=./mock/mock_token_test.go github.com/elangreza14/advance-todo/internal/domain TokenRepository
