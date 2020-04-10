SERVICE=goffmpeg
TEST_DIRECTORY=...

compose_build:
	docker-compose build
test: compose_build
	GO_ENV=test docker-compose \
		run ${SERVICE} \
		go test -v -run=. ./${TEST_DIRECTORY}
tdd: compose_build
	GO_ENV=test docker-compose \
		run ${SERVICE} \
		CompileDaemon \
		-build="go build ./${TEST_DIRECTORY}" \
		-command="go test -failfast -v -run=. ./${TEST_DIRECTORY}" \
		-log-prefix="false"
cover:
	docker-compose run ${SERVICE} \
		go test -coverprofile cover.out ./${TEST_DIRECTORY}
cover-html: cover
	go tool cover -html=cover.out -o cover.html
	open cover.html
