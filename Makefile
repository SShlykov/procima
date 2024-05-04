dev:
	docker compose -f docker-compose.dev.yml up --build

bench:
	docker run --rm -v .:/app -w /app jordi/ab -T application/json -p post.json -k -c 100 -n 500 http://192.168.0.18:8080/api/v1/images:upload
