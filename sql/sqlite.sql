CREATE TABLE "user"
(
    "id"               INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "public_key"       TEXT    NOT NULL,
    "public_key_hash"  TEXT    NOT NULL,
    "private_key"      TEXT    NOT NULL DEFAULT '',
    "private_key_hash" TEXT    NOT NULL DEFAULT '',
    "created_at"       DATE    NOT NULL,
    "updated_at"       DATE    NOT NULL,
    CONSTRAINT "uniq_public_key_hash" UNIQUE ("public_key_hash" ASC)
);

CREATE TABLE "data"
(
    "id"         INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "data_type"  TEXT    NOT NULL,
    "data"       TEXT    NOT NULL,
    "data_hash"  TEXT    NOT NULL,
    "created_at" DATE    NOT NULL,
    "updated_at" DATE    NOT NULL,
    CONSTRAINT "uniq_data_hash" UNIQUE ("data_hash" ASC)
);

CREATE TABLE "data_create"
(
    "id"               INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "data_hash"        TEXT    NOT NULL,
    "public_key_hash"  TEXT    NOT NULL,
    "create_sign"      TEXT    NOT NULL,
    "create_sign_hash" TEXT    NOT NULL,
    "created_at"       DATE    NOT NULL,
    "updated_at"       DATE    NOT NULL
);

CREATE TABLE "data_publish"
(
    "id"                INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "parent_hash"       TEXT    NOT NULL,
    "create_sign_hash"  TEXT    NOT NULL,
    "public_key_hash"   TEXT    NOT NULL,
    "publish_sign"      TEXT    NOT NULL,
    "publish_sign_hash" TEXT    NOT NULL,
    "created_at"        DATE    NOT NULL,
    "updated_at"        DATE    NOT NULL
);

CREATE INDEX "idx_parent_hash_publish_at"
    ON "data_publish" (
                       "parent_hash" ASC,
                       "created_at" ASC
        );