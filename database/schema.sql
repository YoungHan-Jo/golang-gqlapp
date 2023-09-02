-- gen_random_bytes()の使うため拡張機能追加
CREATE EXTENSION pgcrypto;

-- 更新日(updated_at)自動生成 function
CREATE OR REPLACE FUNCTION update_time()
RETURNS TRIGGER AS
'
BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
END;
' LANGUAGE plpgsql;


-- comapny table
CREATE TABLE IF NOT EXISTS "companies" (
	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
	"name" VARCHAR(255) NOT NULL,
	"representative_name" VARCHAR(255) NOT NULL,
	"status" VARCHAR(255) NOT NULL,
	"icon_img_url" VARCHAR(2048) NULL DEFAULT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id")
);
comment on column companies.id IS '会社ID';
comment on column companies.name IS '会社名';
comment on column companies.representative_name IS '代表者名';
comment on column companies.status IS '会社ステータス';
comment on column companies.icon_img_url IS 'アイコン画像URL';
comment on column companies.created_at IS 'データ作成日';
comment on column companies.updated_at IS 'データ更新日';

CREATE TRIGGER update_trigger
BEFORE UPDATE ON companies
FOR EACH ROW
EXECUTE PROCEDURE update_time();

CREATE INDEX IF NOT EXISTS "companies_status_idx"
ON "companies" USING BTREE ("status");


-- hearing_project table
CREATE TABLE IF NOT EXISTS "hearing_projects" (
	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
	"company_id" BYTEA NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"plan" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_company_id" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
comment on column hearing_projects.id IS 'ヒアリングプロジェクトID';
comment on column hearing_projects.company_id IS '会社ID';
comment on column hearing_projects.name IS 'ヒアリングプロジェクト名';
comment on column hearing_projects.plan IS 'サブスクリプションのプラン';
comment on column hearing_projects.created_at IS 'データ作成日';
comment on column hearing_projects.updated_at IS 'データ更新日';

CREATE TRIGGER update_trigger
BEFORE UPDATE ON hearing_projects
FOR EACH ROW
EXECUTE PROCEDURE update_time();


-- users table
CREATE TABLE IF NOT EXISTS "users" (
	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
	"company_id" BYTEA NOT NULL,
	"status" VARCHAR(255) NOT NULL,
	"type" VARCHAR(255) NOT NULL,
	"first_name" VARCHAR(255) NOT NULL,
	"furigana_first_name" VARCHAR(255) NULL DEFAULT NULL,
	"last_name" VARCHAR(255) NOT NULL,
	"furigana_last_name" VARCHAR(255) NULL DEFAULT NULL,
	"middle_name" VARCHAR(255) NULL DEFAULT NULL,
	"email" VARCHAR(255) NOT NULL,
	"icon_img_url" VARCHAR(2048) NOT NULL,
	"industry" VARCHAR(255) NULL DEFAULT NULL,
	"department" VARCHAR(255) NULL DEFAULT NULL,
	"position" VARCHAR(255) NULL DEFAULT NULL,
	"occupation" VARCHAR(255) NULL DEFAULT NULL,
	"remark" VARCHAR(255) NULL DEFAULT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	CONSTRAINT "fk_company_id" FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
comment on column users.id IS 'ユーザID';
comment on column users.company_id IS '会社ID';
comment on column users.status IS 'ユーザステータス';
comment on column users.type IS 'ユーザタイプ（有償・無償）';
comment on column users.first_name IS '名';
comment on column users.furigana_first_name IS 'ふりがな名';
comment on column users.last_name IS '姓';
comment on column users.furigana_last_name IS 'ふりがな姓';
comment on column users.middle_name IS 'ミドルネーム';
comment on column users.email IS 'emailアドレス';
comment on column users.icon_img_url IS 'アイコン画像URL';
comment on column users.industry IS '業種';
comment on column users.department IS '所属部門';
comment on column users.position IS '役職';
comment on column users.occupation IS '職種';
comment on column users.remark IS '備考';
comment on column users.created_at IS 'データ作成日';
comment on column users.updated_at IS 'データ更新日';

CREATE TRIGGER update_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_time();

CREATE INDEX IF NOT EXISTS "users_status_idx"
ON "users" USING BTREE ("status");