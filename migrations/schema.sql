CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT PRIMARY KEY
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"email" TEXT NOT NULL,
"password_hash" TEXT NOT NULL,
"is_superuser" bool,
"active" bool,
"public" bool,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "goals" (
"id" TEXT PRIMARY KEY,
"user_id" user NOT NULL,
"title" TEXT NOT NULL,
"description" TEXT,
"completion_factor" decimal,
"target_date" DATETIME,
"priority" INTEGER,
"tags" varchar[] NOT NULL,
"active" bool,
"public" bool,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "milestones" (
"id" TEXT PRIMARY KEY,
"goal_id" goal NOT NULL,
"title" TEXT NOT NULL,
"description" TEXT,
"contribution_factor" decimal,
"completion_factor" decimal,
"target_date" DATETIME,
"priority" INTEGER,
"tags" varchar[] NOT NULL,
"active" bool,
"public" bool,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "tasks" (
"id" TEXT PRIMARY KEY,
"goal_id" goal NOT NULL,
"milestone_id" milestone NOT NULL,
"title" TEXT NOT NULL,
"description" TEXT,
"contribution_factor" decimal,
"completion_factor" decimal,
"target_date" DATETIME,
"priority" INTEGER,
"tags" varchar[] NOT NULL,
"active" bool,
"public" bool,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
