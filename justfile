[private]
@default:
    just --list

[private]
@clear:
    clear

@dev:
    hugo server -D

@clean:
    echo "🧹 Cleaning build artifacts and caches..."
    rm -rf public node_modules
    pnpm install

@docker-compose-dev:
    docker compose -f compose-dev.yml up --build
