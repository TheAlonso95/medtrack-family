# Database Migrations

This directory contains database migration files for the application.

## Migration File Format

Migration files follow the format:

```
{version}_{name}.up.sql   - For applying migrations
{version}_{name}.down.sql - For rolling back migrations
```

## Creating Migrations

To create a new migration:

```bash
make migrate-create name=migration_name
```

This will create two new files:
- `{timestamp}_migration_name.up.sql`
- `{timestamp}_migration_name.down.sql`

## Running Migrations

To apply all pending migrations:

```bash
make migrate-up
```

To rollback the last migration:

```bash
make migrate-down
```

To check the current migration status:

```bash
make migrate-status
```

## Migration Tool

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations.

To install the migration tool:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```