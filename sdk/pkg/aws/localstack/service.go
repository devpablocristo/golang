package sdkawslocal

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/devpablocristo/golang/sdk/pkg/aws/localstack/ports"
)

var (
	instance  ports.Service
	once      sync.Once
	initError error
)

type service struct {
	config ports.Config
	awsCfg aws.Config
}

func newService(c ports.Config) (ports.Service, error) {
	once.Do(func() {
		svc := &service{config: c}
		initError = svc.Connect()
		if initError != nil {
			instance = nil
		} else {
			instance = svc
		}
	})
	return instance, initError
}

func (s *service) Connect() error {
	// Configurar el SDK de AWS para usar LocalStack
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s.config.GetAWSRegion()),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			s.config.GetAWSAccessKeyID(), s.config.GetAWSSecretAccessKey(), "",
		)),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: s.config.GetLocalStackEndpoint(),
				// Puedes agregar otros ajustes si es necesario
			}, nil
		})),
	)
	if err != nil {
		return err
	}
	s.awsCfg = awsCfg
	return nil
}

func (s *service) GetAWSCfg() aws.Config {
	return s.awsCfg
}
