module mithril

go 1.13

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aliyun/aliyun-oss-go-sdk v2.0.5+incompatible
	github.com/appleboy/gin-jwt/v2 v2.6.3
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/swaggo/swag v1.6.7
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	google.golang.org/grpc v1.34.0
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.2.7
)

replace google.golang.org/grpc v1.34.0 => google.golang.org/grpc v1.26.0
