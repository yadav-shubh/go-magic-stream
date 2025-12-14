docker-setup:
	docker-compose -f database-setup/docker-compose.yaml up

docker-down:
	docker-compose -f database-setup/docker-compose.yaml down -v

# Run client and check if any process is running on 3000 first kill and then run
client-run:
	@echo "Checking if port 3000 is in use..."
	@PID=$$(lsof -t -i:3000); \
	if [ $$PID ]; then \
		echo "Port 3000 is in use by PID $$PID. Killing..."; \
		kill -9 $$PID; \
	else \
		echo "Port 3000 is free."; \
	fi
	cd client && npm run dev


# Run server and check if any process is running on 8080 first kill and then run
server-run:
	@echo "Checking if port 8080 is in use..."
	@PID=$$(lsof -t -i:8080); \
	if [ $$PID ]; then \
		echo "Port 8080 is in use by PID $$PID. Killing..."; \
		kill -9 $$PID; \
	else \
		echo "Port 8080 is free."; \
	fi
	cd server && go run main.go