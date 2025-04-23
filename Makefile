# Variable to frontend url
FRONTEND_URL := http://localhost:8080/signup

run_dev:
	air

run_container:
	docker start 10d7c2aac07f

open_frontend:
	open $(FRONTEND_URL)
