# Database Setup Guide

## Quick Fix for "database nftgenie does not exist"

### Option 1: Using pgAdmin (Recommended)
1. Open pgAdmin
2. Connect to your PostgreSQL server
3. Right-click on "Databases" → "Create" → "Database"
4. Name: `nftgenie`
5. Click "Save"
6. Right-click on the new `nftgenie` database → "Query Tool"
7. Open and run `backend/database/schema.sql`

### Option 2: Using Command Line (if psql is installed)
```bash
# Create database
createdb -U postgres nftgenie

# Apply schema
psql -U postgres -d nftgenie -f backend/database/schema.sql
```

### Option 3: Using SQL Client
1. Connect to PostgreSQL with any SQL client
2. Run: `CREATE DATABASE nftgenie;`
3. Connect to the new database
4. Run the contents of `backend/database/schema.sql`

## Verify Setup
After creating the database, start the backend:
```bash
cd backend
go run main.go
```

You should see: "Connected to database successfully"