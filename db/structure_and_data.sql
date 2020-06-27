/*
 Navicat Premium Data Transfer

 Source Server         : Postgress Local
 Source Server Type    : PostgreSQL
 Source Server Version : 100012
 Source Host           : localhost:5432
 Source Catalog        : crud_multiple_transport
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100012
 File Encoding         : 65001

 Date: 27/06/2020 12:40:13
*/


-- ----------------------------
-- Table structure for gorp_migrations
-- ----------------------------
DROP TABLE IF EXISTS "public"."gorp_migrations";
CREATE TABLE "public"."gorp_migrations" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "applied_at" timestamptz(6)
)
;

-- ----------------------------
-- Records of gorp_migrations
-- ----------------------------
INSERT INTO "public"."gorp_migrations" VALUES ('20200622015928-create_table_users.sql', '2020-06-22 02:06:42.620779+07');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "full_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(50) COLLATE "pg_catalog"."default",
  "password" varchar(128) COLLATE "pg_catalog"."default",
  "mobile_phone" varchar(20) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES ('fa4e391f-3928-4a4b-8204-e0b2dbfde7a5', 'syaikhul hadi', 'my.ant2008@gmail.com', '$2a$04$1jqbY1zgl1pse0WaMJ6afu64lX1QWDGwKuRe08um4WExM8Gk29ncC', '+628814388476', '2020-06-23 17:05:47', '2020-06-23 17:11:03', NULL);
INSERT INTO "public"."users" VALUES ('9fb17677-515b-4c6b-a0af-0fca5f71c969', 'Ferry Anggriawan', 'ferrya1997@gmail.com', '$2a$04$unhQ.9t5gtSyWjnq6oPp..RTfJB2JzOQP8BczSLbKkGF6OuBui4iO', '+6281238564789', '2020-06-23 17:11:40', '2020-06-23 17:11:40', NULL);
INSERT INTO "public"."users" VALUES ('01ce1e75-4fef-4c33-97a8-137e99790c39', 'syaikhul hadi', 'filosteam@gmail.com', '$2a$04$FgLFwLdT/ANpQ4E7ARIIpexIM5tgBHRcgJ.hs0P5iDkd.7LN2ut0m', '+6281238564900', '2020-06-23 17:07:20', '2020-06-23 17:15:32', '2020-06-23 17:15:32');
INSERT INTO "public"."users" VALUES ('51f3f269-9102-42aa-b692-529b924b0a3d', 'Muta Avivi', 'muta.avivi@gmail.com', '$2a$04$JUZ8Om3IkC0tWci8PQ0V/.zlDp6DQQ1Mg221SRwMGxkTfny14ex8C', '+62848778474', '2020-06-23 18:10:06', '2020-06-23 18:10:06', NULL);
INSERT INTO "public"."users" VALUES ('7ec1c85e-a4d4-4fd4-a6bf-e6ee69d3d824', 'Ferry', 'ff@gmail.com', '$2a$04$bggnsnqvhuhVX3wY/ipd2.okHlWxuV4GoTLmWT5fp1lR3sUyVE6K6', '+6281238564', '2020-06-27 03:30:24', '2020-06-27 03:30:24', NULL);
INSERT INTO "public"."users" VALUES ('f1ec534d-e47a-4b3b-bc74-b4304878d434', 'sss', 'ssss@gmail.com', '$2a$04$qO0mzkUfAeb1IKcy4OAnJ.NsZ6YlP0qqBIMPzfZIM9Rqo8GXq3Kpm', '+628123854334', '2020-06-27 03:31:48', '2020-06-27 03:31:48', NULL);
INSERT INTO "public"."users" VALUES ('0af332ff-f78a-4903-b0a6-a5b0a080b1af', 'vivi hadi', 'vivi@gmail.com', '$2a$04$1eaYVioCP8PoAKer9VOyD.GqOXYlNVtA8yST0Sol5cuA6jM.abErO', '+62878678889', '2020-06-27 04:36:22', '2020-06-27 04:53:17', '2020-06-27 04:53:17');

-- ----------------------------
-- Function structure for uuid_generate_v1
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v1mc
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1mc"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1mc"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1mc'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v3
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v3"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v3"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v3'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v4
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v4"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v4"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v4'
  LANGUAGE c VOLATILE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_generate_v5
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v5"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v5"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v5'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_nil
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_nil"();
CREATE OR REPLACE FUNCTION "public"."uuid_nil"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_nil'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_dns
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_dns"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_dns"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_dns'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_oid
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_oid"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_oid"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_oid'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_url
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_url"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_url"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_url'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Function structure for uuid_ns_x500
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_x500"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_x500"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_x500'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;

-- ----------------------------
-- Primary Key structure for table gorp_migrations
-- ----------------------------
ALTER TABLE "public"."gorp_migrations" ADD CONSTRAINT "gorp_migrations_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
