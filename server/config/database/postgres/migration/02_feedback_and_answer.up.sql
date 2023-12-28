CREATE TABLE IF NOT EXISTS "tb_answers" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "answered_by" UUID NOT NULL,
  "content" TEXT NOT NULL,
  "answered_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT "fk_answered_by"
  FOREIGN KEY ("answered_by")
  REFERENCES "tb_employees" ("registry")
);

CREATE TABLE IF NOT EXISTS "tb_feedbacks" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "employee_registry" UUID NOT NULL,
  "content" TEXT NOT NULL,
  "answered" BOOLEAN DEFAULT FALSE,
  "answer_id" UUID UNIQUE,
  "sent_at" TIMESTAMP DEFAULT NOW(),
  CONSTRAINT "fk_employee_registry"
  FOREIGN KEY ("employee_registry")
  REFERENCES "tb_employees" ("registry"),
  CONSTRAINT "fk_answer_id"
  FOREIGN KEY ("answer_id")
  REFERENCES "tb_answers" ("id")
);