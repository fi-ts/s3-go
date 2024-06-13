package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	v1 "github.com/fi-ts/s3-go/pkg/apis/v1"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type S3Client interface {
	User() v1.S3UserServiceClient
	Partition() v1.S3PartitionServiceClient
	Health() healthv1.HealthClient
	Close() error
}

type client struct {
	conn *grpc.ClientConn
}

func NewClient(hostname string, port int, certFile string, keyFile string, caFile string) (S3Client, error) {
	address := fmt.Sprintf("%s:%d", hostname, port)

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to load system credentials: %w", err)
	}

	if caFile != "" {
		ca, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("could not read ca certificate: %w", err)
		}

		ok := certPool.AppendCertsFromPEM(ca)
		if !ok {
			return nil, fmt.Errorf("failed to append ca certs: %s", caFile)
		}
	}

	clientCertificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %w", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   hostname,
		Certificates: []tls.Certificate{clientCertificate},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}

	return client{
		conn: conn,
	}, nil
}

// Close the underlying connection
func (c client) Close() error {
	return c.conn.Close()
}

// User is the root accessor for s3 user related functions
func (c client) User() v1.S3UserServiceClient {
	return v1.NewS3UserServiceClient(c.conn)
}

// Partition is the root accessor for s3 partition related functions
func (c client) Partition() v1.S3PartitionServiceClient {
	return v1.NewS3PartitionServiceClient(c.conn)
}

// Health is the root accessor for health related functions
func (c client) Health() healthv1.HealthClient {
	return healthv1.NewHealthClient(c.conn)
}
