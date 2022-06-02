package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	v1 "github.com/fi-ts/s3-go/pkg/apis/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
)

type S3ControllerClient interface {
	User() v1.S3ControllerUserServiceClient
	Health() healthv1.HealthClient
	Close() error
}

type client struct {
	conn *grpc.ClientConn
}

func NewClient(ctx context.Context, hostname string, port int, cert []byte, certKey []byte, ca []byte) (S3ControllerClient, error) {
	address := fmt.Sprintf("%s:%d", hostname, port)

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to load system credentials: %w", err)
	}

	if string(ca) != "" {
		ok := certPool.AppendCertsFromPEM(ca)
		if !ok {
			return nil, fmt.Errorf("failed to append ca certs: %s", ca)
		}
	}

	clientCertificate, err := tls.X509KeyPair(cert, certKey)
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
		grpc.WithBlock(),
	}

	conn, err := grpc.DialContext(ctx, address, opts...)
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
func (c client) User() v1.S3ControllerUserServiceClient {
	return v1.NewS3ControllerUserServiceClient(c.conn)
}

// Health is the root accessor for health related functions
func (c client) Health() healthv1.HealthClient {
	return healthv1.NewHealthClient(c.conn)
}
