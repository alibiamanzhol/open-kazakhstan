# Open Kazakhstan

A compact Go toolkit for exploring Kazakhstan-oriented open-data catalogs from the command line. It is designed around typed data models, interchangeable local/HTTP providers, deterministic output and zero third-party runtime dependencies.

## What it does

```bash
go run ./cmd/kzdata list
go run ./cmd/kzdata list --query "transport"
go run ./cmd/kzdata show demo-001
go run ./cmd/kzdata stats
go run ./cmd/kzdata export --format json --out datasets.json
```

The default provider is an embedded deterministic fixture, so the project works offline and is easy to test. To connect a compatible catalog endpoint:

```bash
KZDATA_URL=https://example.kz/api/datasets go run ./cmd/kzdata list
```

The endpoint may return either a JSON array of datasets or an object containing a `datasets` array.

## Architecture

```text
cmd/kzdata/       command routing and flags
internal/client/  local and HTTP data providers
internal/model/   domain types
internal/render/  table, JSON and CSV output
```

The command layer stays deliberately small. Provider-specific behavior is isolated from rendering, and output code can be tested without network access.

## Quality checks

```bash
go test ./...
go vet ./...
```

Both checks run on every push and pull request through GitHub Actions.

## License

MIT
