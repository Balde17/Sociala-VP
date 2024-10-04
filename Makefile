stash:
	git stash

pull: stash
	git pull

drop: stash
	git stash drop

apply: pull
	git stash apply	

push: pull apply
	git add .
	git commit -m "$(m)"
	git push

gen-key:
	@echo "SECRET_KEY=$$(openssl rand -base64 $(base))" >> backend/app/.env

create-handler:
	@mkdir backend/app/internals/handlers/$(handler)
	@echo 'package handlers\n\nimport "net/http"\n\nfunc $(handler)(w http.ResponseWriter, r *http.Request) {\n\n}' > backend/app/internals/handlers/$(handler)/$(handler).go

rm-handler:
	rm -rf backend/app/internals/handlers/$(handler)

start:
	cd backend/app/ && go run .	

teamCanada:
	cd frontend && npm run dev

execMigrate:
	@sqlite3 backend/app/social.db < $(filename)

install:
	@cd frontend && npm install	

deploy:
	cd frontend && vercel

exec:
	@sqlite3 backend/app/social.db < database/exec.sql