.PHONY: protolint
protolint:
	buf format -w metalstack/io/s3/api/v1
	buf lint -v

.PHONY: protoc
protoc: protolint
	buf generate -v