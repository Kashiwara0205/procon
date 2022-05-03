in:
	docker exec -ti procon-app-1 /bin/bash

go_test:
	docker exec -ti procon-app-1 bash -c "cd procon; go test -v ./..."

go_fmt:
	docker exec -ti procon-app-1 bash -c "cd procon; go fmt ./..."
	git add -A; 
	git commit -m "fmt適用"