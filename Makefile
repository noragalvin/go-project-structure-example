
hello:
	echo "Hello World"

ps:
	docker-compose -f deployments/docker-compose.yml --env-file .env ps

stop:
	docker-compose -f deployments/docker-compose.yml --env-file .env stop

rm:
	docker-compose -f deployments/docker-compose.yml --env-file .env rm -f

down:
	docker-compose -f deployments/docker-compose.yml --env-file .env down --remove-orphans

### Woocommerce
woo-watch:
	docker-compose -f deployments/docker-compose.yml --env-file .env up woo_app

woo-exec-app:
	docker-compose -f deployments/docker-compose.yml --env-file .env exec woo_app bash

woo-exec-database:
	docker-compose -f deployments/docker-compose.yml --env-file .env exec database bash

woo-seeding:
	go run cmd/seeding/woocommerce/main.go

woo-build:
	# TODO

### End Woocommerce

### Shopify

shopify-watch:
	docker-compose -f deployments/docker-compose.yml --env-file .env up shopify_app

shopify-exec-app:
	docker-compose -f deployments/docker-compose.yml --env-file .env exec shopify_app bash

shopify-exec-database:
	docker-compose -f deployments/docker-compose.yml --env-file .env exec database bash

shopify-seeding:
	go run cmd/seeding/shopify/main.go

shopify-build:
	# TODO

### End Shopify
