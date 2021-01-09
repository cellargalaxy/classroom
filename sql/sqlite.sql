CREATE TABLE "user" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "public_key" TEXT NOT NULL,
    "public_key_hash" TEXT NOT NULL,
    "private_key" TEXT,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NOT NULL,
    "deleted_at" DATE DEFAULT NULL,
    CONSTRAINT "uniq_public_key_hash" UNIQUE ("public_key_hash" ASC)
);

CREATE TABLE "message" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "user_hash" TEXT NOT NULL,
    "parent_hash" TEXT NOT NULL,
    "data" TEXT NOT NULL,
    "data_type" TEXT NOT NULL,
    "hash" TEXT NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NOT NULL,
    "deleted_at" DATE DEFAULT NULL
);

CREATE INDEX "idx_hash"
ON "message" (
    "hash" ASC,
    "created_at" DESC
);

CREATE INDEX "idx_parent_hash"
ON "message" (
    "parent_hash" ASC,
    "created_at" DESC
);

CREATE INDEX "idx_user_hash"
ON "message" (
    "user_hash" ASC,
    "created_at" DESC
);