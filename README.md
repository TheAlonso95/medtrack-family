# Monorepo Project

This repository contains a monorepo with the following components:

- `/backend`: Go backend service
- `/web`: Next.js frontend application
- `/mobile`: Kotlin Multiplatform Mobile application targeting Android and iOS

## Development

### Backend (Go)
```bash
cd backend
go test ./...
```

### Web (Next.js)
```bash
cd web
yarn install
yarn dev
```

### Mobile (Kotlin Multiplatform)
```bash
cd mobile
./gradlew build
```

## CI/CD

This repository uses GitHub Actions for continuous integration. The workflow runs:
- Go formatting checks and tests
- Next.js build
- Kotlin Multiplatform build

See `.github/workflows/ci.yml` for details.