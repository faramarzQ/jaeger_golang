up:
	docker-compose up

run-app:
	( JAEGER_SERVICE_NAME=Jaeger_golang JAEGER_SAMPLER_PARAM=1 go run etc/main.go )