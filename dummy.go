package dummy

import
(
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DataDog/datadog-go"
	"github.com/aws/aws-sdk-go"
	"github.com/buger/jsonparser"
	"github.com/creasty/defaults"
	"github.com/deliveryhero/pd-fluid-go-sdk/fluid"
	"github.com/deliveryhero/pd-go-kit"
	"github.com/deliveryhero/pd-oneview-go-toolkit"
	"github.com/deliveryhero/pd-protobuf-go/customer-communications"
	"github.com/deliveryhero/pd-protobuf-go/pablo"
	"github.com/deliveryhero/pd-sts-go-sdk"
	"github.com/dgraph-io/ristretto"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/golang/protobuf"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify"
	"github.com/wI2L/jsondiff"
	"golang.org/x/tools"
	"google.golang.org/grpc"
	"google.golang.org/protobuf"
	"gopkg.in/DataDog/dd-trace-go.v1"
	"gopkg.in/yaml.v2"
	"gopkg.in/yaml.v3"
)

func main() {
	return
}