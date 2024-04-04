CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT PRIMARY KEY
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"email" TEXT NOT NULL,
"password" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "goals" (
"id" TEXT PRIMARY KEY,
"user" user NOT NULL,
"title" TEXT NOT NULL,
"description" TEXT,
"target_date" DATETIME,
"priority" INTEGER,
"tags" varchar[] NOT NULL,
"active" bool,
"public" bool,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
