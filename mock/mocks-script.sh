#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

# lib
mockgen -package=mock -destination=mock/mock_ResponseWriter.go -source=/usr/local/go/src/net/http/server.go

# controller
mockgen -package=mock -destination=mock/mock_HealthCheckController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IHealthCheckController
mockgen -package=mock -destination=mock/mock_AuthController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IAuthController
mockgen -package=mock -destination=mock/mock_SegmentController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller ISegmentController
mockgen -package=mock -destination=mock/mock_BillController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IBillController
mockgen -package=mock -destination=mock/mock_PaylaterHistory.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IPaylaterHistoryController
mockgen -package=mock -destination=mock/mock_PaylaterRepay.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IPaylaterRepayController
mockgen -package=mock -destination=mock/mock_ScenarioCronController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IScenarioCronController
mockgen -package=mock -destination=mock/mock_CallbackController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller ICallbackController
mockgen -package=mock -destination=mock/mock_MonkeyPatchController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IMonkeyPatchController
mockgen -package=mock -destination=mock/mock_NotificationCronController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller INotificationCronController
mockgen -package=mock -destination=mock/mock_ScenarioController.go gitlab.mapan.io/mapan-pulsa/service/paylater/controller IScenarioController

# service
mockgen -package=mock -destination=mock/mock_HealthCheckService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IHealthCheckService
mockgen -package=mock -destination=mock/mock_AuthService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IAuthService
mockgen -package=mock -destination=mock/mock_SegmentService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service ISegmentService
mockgen -package=mock -destination=mock/mock_PriceService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IPriceService
mockgen -package=mock -destination=mock/mock_BillService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IBillService
mockgen -package=mock -destination=mock/mock_ChargeService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IChargeService
mockgen -package=mock -destination=mock/mock_RefundService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IRefundService
mockgen -package=mock -destination=mock/mock_PaylaterRepayService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IPaylaterRepayService
mockgen -package=mock -destination=mock/mock_TransactionService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service ITransactionService
mockgen -package=mock -destination=mock/mock_ScenarioService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service IScenarioService
mockgen -package=mock -destination=mock/mock_NotificationService.go gitlab.mapan.io/mapan-pulsa/service/paylater/service INotificationService

# repository
mockgen -package=mock -destination=mock/mock_HealthCheckRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IHealthCheckRepository
mockgen -package=mock -destination=mock/mock_AuthRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IAuthRepository
mockgen -package=mock -destination=mock/mock_SegmentRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository ISegmentRepository
mockgen -package=mock -destination=mock/mock_PriceRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IPriceRepository
mockgen -package=mock -destination=mock/mock_PaylaterTransactionRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IPaylaterTransactionRepository
mockgen -package=mock -destination=mock/mock_PaylaterRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IPaylaterRepository
mockgen -package=mock -destination=mock/mock_ScenarioRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IScenarioRepository
mockgen -package=mock -destination=mock/mock_IdempotencyRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IIdempotencyRepository
mockgen -package=mock -destination=mock/mock_PaylaterRepayRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IPaylaterRepayRepository
mockgen -package=mock -destination=mock/mock_PaylaterHistoryRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IPaylaterHistoryRepository
mockgen -package=mock -destination=mock/mock_TransactionRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository ITransactionRepository
mockgen -package=mock -destination=mock/mock_PaylaterWithPaylaterRepayTXRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IPaylaterWithPaylaterRepayTXRepository
mockgen -package=mock -destination=mock/mock_NotificationRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository INotificationRepository
mockgen -package=mock -destination=mock/mock_PaylaterWithPaylaterTransactionTXRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IPaylaterWithPaylaterTransactionTXRepository
mockgen -package=mock -destination=mock/mock_ChargeRepository.go gitlab.mapan.io/mapan-pulsa/service/paylater/repository IChargeRepository

# requester
mockgen -package=mock -destination=mock/requester/mock_MSARequester.go gitlab.mapan.io/mapan-pulsa/service/paylater/requester/mapansocial IMapanSocial
mockgen -package=mock -destination=mock/requester/mock_Midtrans.go gitlab.mapan.io/mapan-pulsa/service/paylater/requester/midtrans IMidtrans

# Logger
mockgen -package=mock -destination=mock/mock_Logger.go gitlab.mapan.io/mapan-go-template/go-logger/logger ILogger

# HTTP Client
mockgen -package=mock -destination=mock/mock_HTTPClient.go gitlab.mapan.io/mapan-go-template/go-http-client/client IClient

# Database
mockgen -package=mock -destination=mock/mock_Transaction.go gitlab.mapan.io/mapan-go-template/go-db-sql ITransaction
mockgen -package=mock -destination=mock/mock_Database.go gitlab.mapan.io/mapan-go-template/go-db-sql IDatabase
mockgen -package=mock -destination=mock/mock_Row.go gitlab.mapan.io/mapan-go-template/go-db-sql IRow
mockgen -package=mock -destination=mock/mock_Rows.go gitlab.mapan.io/mapan-go-template/go-db-sql IRows
mockgen -package=mock -destination=mock/mock_Result.go gitlab.mapan.io/mapan-go-template/go-db-sql IResult

# Redis
mockgen -package=mock -destination=mock/mock_Redis.go gitlab.mapan.io/mapan-go-template/go-redis IRedis

# Kafka
mockgen -package=mock -destination=mock/mock_KafkaProducer.go gitlab.mapan.io/mapan-go-template/go-kafka Producer