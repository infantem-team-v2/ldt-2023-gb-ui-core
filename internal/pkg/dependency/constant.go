package dependency

import (
	"gb-ui-core/config"
	authRepo "gb-ui-core/internal/auth/repository"
	authUC "gb-ui-core/internal/auth/usecase"
	calculatorUC "gb-ui-core/internal/calculator/usecase"
	mdwHttp "gb-ui-core/internal/pkg/middleware/delivery/http"
	uiRepo "gb-ui-core/internal/ui/repository"
	"gb-ui-core/pkg/damqp/kafka"
	"gb-ui-core/pkg/damqp/rabbit"
	"gb-ui-core/pkg/terrors"
	"gb-ui-core/pkg/thttp"
	"gb-ui-core/pkg/thttp/server"
	"gb-ui-core/pkg/tlogger"
	"gb-ui-core/pkg/tsecure"
	tstorageCache "gb-ui-core/pkg/tstorage/cache"
	tstorageNonRelational "gb-ui-core/pkg/tstorage/nonrelational"
	tstorageRelational "gb-ui-core/pkg/tstorage/relational"
	"github.com/sarulabs/di"
)

var dependencyMap = map[string]func(ctn di.Container) (interface{}, error){
	"config": config.BuildConfig,

	"fernet": tsecure.BuildFernetEncryptor,

	"postgres": tstorageRelational.BuildPostgres,
	"redis":    tstorageCache.BuildRedis,
	"mongo":    tstorageNonRelational.BuildMongoDB,

	"httpClient": thttp.BuildHttpClient,

	"logger": tlogger.BuildLogger,

	"rabbit": rabbit.BuildRabbitMQ,
	"kafka":  kafka.BuildKafka,

	"authUC":   authUC.BuildAuthUsecase,
	"authRepo": authRepo.BuildPostgresRepository,

	"middleware":        mdwHttp.BuildMiddlewareManager,
	"errorHandler":      terrors.BuildErrorHandler,
	"stacktraceHandler": terrors.BuildStacktraceHandler,

	"uiMongoRepo":    uiRepo.BuildMongoRepository,
	"uiPostgresRepo": uiRepo.BuildPostgresRepository,

	"calcUC": calculatorUC.BuildCalculatorUseCase,

	"app": server.BuildFiberApp,
}

const TagDI = "di"
