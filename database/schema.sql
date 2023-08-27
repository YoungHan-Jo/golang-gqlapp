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


-- -- hearing_project_member table
-- CREATE TABLE IF NOT EXISTS "hearing_project_members" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"user_id" BYTEA NOT NULL,
-- 	"hearing_project_id" BYTEA NOT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
-- 	CONSTRAINT "fk_hearing_project_id" FOREIGN KEY ("hearing_project_id") REFERENCES "hearing_projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column hearing_project_members.id IS 'ヒアリングプロジェクトメンバーID';
-- comment on column hearing_project_members.user_id IS 'ユーザID';
-- comment on column hearing_project_members.hearing_project_id IS 'ヒアリングプロジェクトID';
-- comment on column hearing_project_members.created_at IS 'データ作成日';
-- comment on column hearing_project_members.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON hearing_project_members
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- hearing table
-- CREATE TABLE IF NOT EXISTS "hearings" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"hearing_project_id" BYTEA NOT NULL,
-- 	"company_name" VARCHAR(255) NOT NULL,
-- 	"status" VARCHAR(255) NOT NULL,
-- 	"staff" VARCHAR(255) NOT NULL,
-- 	"reserved_at" TIMESTAMP NULL DEFAULT NULL,
-- 	"closed_at" TIMESTAMP NULL DEFAULT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_hearing_project_id" FOREIGN KEY ("hearing_project_id") REFERENCES "hearing_projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column hearings.id IS 'ヒアリングID';
-- comment on column hearings.hearing_project_id IS 'ヒアリングプロジェクトID';
-- comment on column hearings.company_name IS '企業名';
-- comment on column hearings.status IS 'ヒアリングステータス';
-- comment on column hearings.staff IS '担当者';
-- comment on column hearings.reserved_at IS '開始予定日時';
-- comment on column hearings.closed_at IS '終了予定日時';
-- comment on column hearings.created_at IS 'データ作成日';
-- comment on column hearings.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON hearings
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();

-- CREATE UNIQUE INDEX IF NOT EXISTS "hearings_status_idx"
-- ON "hearings" USING BTREE ("status");


-- -- payment_user table
-- CREATE TABLE IF NOT EXISTS "payment_users" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"user_id" BYTEA NOT NULL,
-- 	"hearing_project_id" BYTEA NOT NULL,
-- 	"stripe_customer_id" VARCHAR(255) NOT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
-- 	CONSTRAINT "fk_hearing_project_id" FOREIGN KEY ("hearing_project_id") REFERENCES "hearing_projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column payment_users.id IS '決済担当者ID';
-- comment on column payment_users.user_id IS 'ユーザID';
-- comment on column payment_users.hearing_project_id IS 'ヒアリングプロジェクトID';
-- comment on column payment_users.stripe_customer_id IS 'stripe customer ID';
-- comment on column payment_users.created_at IS 'データ作成日';
-- comment on column payment_users.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON payment_users
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- payment_histories
-- CREATE TABLE IF NOT EXISTS "payment_histories" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"hearing_project_id" BYTEA NOT NULL,
-- 	"payment_user_id" BYTEA NOT NULL,
-- 	"payment_type" VARCHAR(255) NOT NULL,
-- 	"amount" INTEGER NOT NULL,
-- 	"stripe_pi_id" VARCHAR(255) NOT NULL,
-- 	"payment_date" TIMESTAMP NOT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_hearing_project_id" FOREIGN KEY ("hearing_project_id") REFERENCES "hearing_projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
-- 	CONSTRAINT "fk_payment_user_id" FOREIGN KEY ("payment_user_id") REFERENCES "payment_users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column payment_histories.id IS '決済履歴ID';
-- comment on column payment_histories.hearing_project_id IS 'ヒアリングプロジェクトID';
-- comment on column payment_histories.payment_user_id IS '決済担当者ID';
-- comment on column payment_histories.payment_type IS '決済タイプ（定期/都度）';
-- comment on column payment_histories.amount IS '決済価格';
-- comment on column payment_histories.stripe_pi_id IS 'stripe payment intent ID';
-- comment on column payment_histories.payment_date IS '決済日';
-- comment on column payment_histories.created_at IS 'データ作成日';
-- comment on column payment_histories.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON payment_histories
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- tickets table
-- CREATE TABLE IF NOT EXISTS "tickets" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"hearing_project_id" BYTEA NOT NULL,
-- 	"payment_history_id" BYTEA NOT NULL,
-- 	"is_purchased" BOOLEAN NOT NULL,
-- 	"is_used" BOOLEAN NOT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_hearing_project_id" FOREIGN KEY ("hearing_project_id") REFERENCES "hearing_projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
-- 	CONSTRAINT "fk_payment_history_id" FOREIGN KEY ("payment_history_id") REFERENCES "payment_histories" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column tickets.id IS 'チケットID';
-- comment on column tickets.hearing_project_id IS 'ヒアリングプロジェクトID';
-- comment on column tickets.payment_history_id IS '決済履歴ID';
-- comment on column tickets.is_purchased IS '購入可否';
-- comment on column tickets.is_used IS '使用可否';
-- comment on column tickets.created_at IS 'データ作成日';
-- comment on column tickets.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON tickets
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- ticket_histories table
-- CREATE TABLE IF NOT EXISTS "ticket_histories" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"ticket_id" BYTEA NOT NULL,
-- 	"hearing_id" BYTEA NOT NULL,
-- 	"used_date" TIMESTAMP NOT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_ticket_id" FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
-- 	CONSTRAINT "fk_hearing_id" FOREIGN KEY ("hearing_id") REFERENCES "hearings" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column ticket_histories.id IS 'チケット使用履歴ID';
-- comment on column ticket_histories.ticket_id IS 'チケットID';
-- comment on column ticket_histories.hearing_id IS 'ヒアリングID';
-- comment on column ticket_histories.used_date IS 'チケット使用日';
-- comment on column ticket_histories.created_at IS 'データ作成日';
-- comment on column ticket_histories.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON ticket_histories
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- project_specific_solution table
-- CREATE TABLE IF NOT EXISTS "project_specific_solutions" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"hearing_project_id" BYTEA NOT NULL,
-- 	"name" VARCHAR(255) NOT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_hearing_project_id" FOREIGN KEY ("hearing_project_id") REFERENCES "hearing_projects" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column project_specific_solutions.id IS '個別ソリューションID';
-- comment on column project_specific_solutions.hearing_project_id IS 'ヒアリングプロジェクトID';
-- comment on column project_specific_solutions.name IS 'ソリューション名';
-- comment on column project_specific_solutions.created_at IS 'データ作成日';
-- comment on column project_specific_solutions.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON project_specific_solutions
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- tasks table
-- CREATE TABLE IF NOT EXISTS "tasks" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"hearing_id" BYTEA NOT NULL,
-- 	"hearing_status" VARCHAR(255) NOT NULL,
-- 	"name" VARCHAR(255) NULL DEFAULT NULL,
-- 	"staff" VARCHAR(255) NULL DEFAULT NULL,
-- 	"category" VARCHAR(255) NULL DEFAULT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_hearing_id" FOREIGN KEY ("hearing_id") REFERENCES "hearings" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column tasks.id IS 'ユーザID';
-- comment on column tasks.hearing_id IS 'ヒアリングID';
-- comment on column tasks.hearing_status IS 'ヒアリング状況';
-- comment on column tasks.name IS '業務名';
-- comment on column tasks.staff IS '担当者';
-- comment on column tasks.category IS 'カテゴリ';
-- comment on column tasks.created_at IS 'データ作成日';
-- comment on column tasks.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON tasks
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();

-- CREATE INDEX IF NOT EXISTS "tasks_hearing_id_idx"
-- ON "tasks" USING BTREE ("hearing_id");

-- CREATE INDEX IF NOT EXISTS "tasks_hearing_status_idx"
-- ON "tasks" USING BTREE ("hearing_status");


-- -- task_process table
-- CREATE TABLE IF NOT EXISTS "task_process" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"task_id" BYTEA NOT NULL ,
-- 	"name" VARCHAR(255) NULL DEFAULT NULL,
-- 	"staff" VARCHAR(255) NULL DEFAULT NULL,
-- 	"kind" VARCHAR(255) NULL DEFAULT NULL,
-- 	"system_name" VARCHAR(255) NULL DEFAULT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_task_id" FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column task_process.id IS '業務プロセスID';
-- comment on column task_process.task_id IS '業務ID';
-- comment on column task_process.name IS '処理名';
-- comment on column task_process.staff IS '担当者';
-- comment on column task_process.kind IS '処理の種類';
-- comment on column task_process.system_name IS 'システム名';
-- comment on column task_process.created_at IS 'データ作成日';
-- comment on column task_process.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON task_process
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- task_efficiency table
-- CREATE TABLE IF NOT EXISTS "task_efficiency" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"task_id" BYTEA NOT NULL,
-- 	"head_count" INTEGER NULL DEFAULT NULL,
-- 	"per_hour" INTEGER NULL DEFAULT NULL,
-- 	"frequency1" INTEGER NULL DEFAULT NULL,
-- 	"frequency2" INTEGER NULL DEFAULT NULL,
-- 	"hour" INTEGER NULL DEFAULT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_task_id" FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column task_efficiency.id IS '費用対効率ID';
-- comment on column task_efficiency.task_id IS '業務ID';
-- comment on column task_efficiency.head_count IS '人数';
-- comment on column task_efficiency.per_hour IS '一人当たりの時給';
-- comment on column task_efficiency.frequency1 IS '頻度1';
-- comment on column task_efficiency.frequency2 IS '頻度2';
-- comment on column task_efficiency.hour IS '時間';
-- comment on column task_efficiency.created_at IS 'データ作成日';
-- comment on column task_efficiency.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON task_efficiency
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();


-- -- task_memo table
-- CREATE TABLE IF NOT EXISTS "task_memos" (
-- 	"id" BYTEA NOT NULL DEFAULT gen_random_bytes(16),
-- 	"task_id" BYTEA NOT NULL,
-- 	"content" TEXT NOT NULL,
-- 	"created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	"updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	PRIMARY KEY ("id"),
-- 	CONSTRAINT "fk_task_id" FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
-- );
-- comment on column task_memos.id IS 'タスクメモID';
-- comment on column task_memos.task_id IS 'タスクID';
-- comment on column task_memos.content IS 'メモ内容';
-- comment on column task_memos.created_at IS 'データ作成日';
-- comment on column task_memos.updated_at IS 'データ更新日';

-- CREATE TRIGGER update_trigger
-- BEFORE UPDATE ON task_memos
-- FOR EACH ROW
-- EXECUTE PROCEDURE update_time();