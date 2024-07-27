package svc

import (
	"context"
	AwsConf "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cedarHH/LandTradingContract/api/contract"
	"github.com/cedarHH/LandTradingContract/api/internal/accountAuth"
	"github.com/cedarHH/LandTradingContract/api/internal/config"
	"github.com/cedarHH/LandTradingContract/model/database"
	"github.com/cedarHH/LandTradingContract/model/objectStorage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

type ServiceContext struct {
	Config      config.Config
	EthClient   *ethclient.Client
	Instance    *contract.Contract
	Conn        *contract.Contract
	AccountAuth *accountAuth.AccountAuth
	LandModel   database.LandModel
	LandBucket  objectStorage.IS3Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	landTableConfig, err := AwsConf.LoadDefaultConfig(context.TODO(), AwsConf.WithRegion(c.DynamoDBConf.Region))
	if err != nil {
		panic("Failed to load UserTable config: " + err.Error())
	}
	s3Config, err := AwsConf.LoadDefaultConfig(context.TODO(), AwsConf.WithRegion(c.S3Conf.Region))
	if err != nil {
		panic("Failed to load S3 config: " + err.Error())
	}

	s3Client := s3.NewFromConfig(s3Config)
	presignClient := s3.NewPresignClient(s3Client)
	landBucket := objectStorage.NewLandModel(s3Client, presignClient, c.S3Conf.Bucket)
	landDynamoDBClient := dynamodb.NewFromConfig(landTableConfig)
	landModel := database.NewLandModel(landDynamoDBClient, c.DynamoDBConf.TableName)

	ganacheURL := os.Getenv("GANACHE_URL")
	client, err := ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	accAuth := accountAuth.NewAccountAuth(client)

	privateKey := os.Getenv("PRIVATE_KEY")
	auth := accAuth.GetAccountAuth(privateKey)

	address, tx, instance, err := contract.DeployContract(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	_, _ = instance, tx

	log.Printf("Contract deployed at address: %s", address.Hex())

	conn, err := contract.NewContract(common.HexToAddress(address.Hex()), client)
	if err != nil {
		log.Fatalf("Failed to create api object: %v", err)
	}

	return &ServiceContext{
		Config:      c,
		EthClient:   client,
		Instance:    instance,
		Conn:        conn,
		AccountAuth: accAuth,
		LandModel:   landModel,
		LandBucket:  landBucket,
	}
}
