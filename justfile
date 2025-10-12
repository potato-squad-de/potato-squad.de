set dotenv-load

[private]
@default:
    just --list

[private]
@clear:
    clear

@clean:
    echo "🧹 Cleaning build artifacts and caches..."
    rm -rf public node_modules tmp dist
    pnpm install

@docker-compose-dev:
    docker compose -f compose-dev.yml up --build

@format:
	go fmt ./...

@lint:
    golangci-lint run ./...

@update:
	go get -u ./...

@dev-site: clear
    hugo server

@build-site:
    hugo

# @dev-web *args: clear
# 	wgo -file .go go mod tidy :: wgo -file .go -file .sql go run -tags='dev' ./cmd/web :: wgo -dir layouts -dir content -dir static -dir assets -xfile assets/jsconfig.json hugo
#

@dev-discord-bot *args: clear
    wgo -file .go go mod tidy :: wgo -file .go -file .sql go run -tags='dev' ./cmd/discord-bot

@build-discord-bot:
    go build -o ./dist/discord-bot ./cmd/discord-bot/
