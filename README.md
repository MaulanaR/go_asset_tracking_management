# Asset Tracking Management

## Getting Started
1. Make sure you have [Go](https://go.dev) installed.
2. Clone the repo
```bash
git clone https://github.com/maulanar/go_asset_tracking_management.git
```
3. Go to the directory and run go mod tidy to add missing requirements and to drop unused requirements
```bash
cd go_asset_tracking_management && go mod tidy
```
3. Setup your .env file
```bash
cp .env-example .env && vi .env
```
4. Start
```bash
go run main.go
```

## Open API Documentation
1. Update your open api documentation
```bash
go run main.go update
```
2. Start
```bash
go run main.go
```
3. Open /api/docs in browser

## Test
1. Make sure you have db with name `main_test.db` with credentials same as DB_XXX
2. Test all with verbose output that lists all of the tests and their results.
```bash
ENV_FILE=$(pwd)/.env go test ./... -v
```
3. Test all with benchmark.
```bash
ENV_FILE=$(pwd)/.env go test ./... -bench=.
```

## Build for production
1. Compile packages and dependencies
```bash
go build -o go_asset_tracking_management main.go
```
2. Setup .env file for production
```bash
cp .env-example .env && vi .env
```
3. Run executable file with systemd, supervisor, pm2 or other process manager
```bash
./go_asset_tracking_management
```