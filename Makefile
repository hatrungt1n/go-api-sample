#================================
#== GOLANG
#================================
GO := @go
# GIN := @gin

goinstall:
	${GO} get .

godev:
	${GO} run main.go

goprod:
	${GO} build -o main .

# gotest:
# 	${GO} test -v

gofmt:
	${GO} fmt ./...