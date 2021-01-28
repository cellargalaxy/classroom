CREATE TABLE "user"
(
    "id"              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "public_key"      TEXT    NOT NULL,
    "public_key_hash" TEXT    NOT NULL,
    "private_key"     TEXT,
    "created_at"      DATE    NOT NULL,
    "updated_at"      DATE    NOT NULL,
    CONSTRAINT "uniq_public_key_hash" UNIQUE ("public_key_hash" ASC)
);

CREATE TABLE "message"
(
    "id"                      INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "data_type"               TEXT    NOT NULL,
    "data"                    TEXT    NOT NULL,
    "data_hash"               TEXT    NOT NULL,
    "create_sign"             TEXT    NOT NULL,
    "create_public_key_hash"  TEXT    NOT NULL,
    "create_sign_hash"        TEXT    NOT NULL,
    "publish_sign"            TEXT    NOT NULL,
    "publish_public_key_hash" TEXT    NOT NULL,
    "publish_sign_hash"       TEXT    NOT NULL,
    "parent_hash"             TEXT    NOT NULL,
    "publish_at"              DATE    NOT NULL,
    "created_at"              DATE    NOT NULL,
    "updated_at"              DATE    NOT NULL
);

CREATE INDEX "idx_data_hash"
    ON "message" (
                  "data_hash" ASC
        );

CREATE INDEX "idx_create_sign_hash"
    ON "message" (
                  "create_sign_hash" ASC
        );

CREATE INDEX "idx_publish_sign_hash"
    ON "message" (
                  "publish_sign_hash" ASC
        );

CREATE INDEX "idx_parent_hash_publish_at"
    ON "message" (
                  "parent_hash" ASC,
                  "publish_at" DESC
        );