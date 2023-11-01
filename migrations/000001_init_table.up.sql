BEGIN;

CREATE TABLE "users" (
                         "id" uuid PRIMARY KEY,
                         "username" varchar(255) NOT NULL,
                         "full_name" varchar(255) NOT NULL,
                         "hash_password" varchar(255),
                         "salt" varchar(255),
                         "created_at" timestamp,
                         "updated_at" timestamp,
                         "deleted_at" timestamp
);

CREATE TABLE "sessions" (
                            "id" uuid PRIMARY KEY,
                            "user_id" uuid NOT NULL,
                            "refresh_token" varchar NOT NULL,
                            "is_blocked" boolean NOT NULL DEFAULT false,
                            "expires_at" timestamp NOT NULL ,
                            "created_at" timestamp,
                            "updated_at" timestamp,
                            "deleted_at" timestamp
);



ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

COMMIT;