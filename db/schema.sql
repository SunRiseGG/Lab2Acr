CREATE TABLE "banking" (
    "id"   SERIAL PRIMARY KEY,
    "balance" DECIMAL NOT NULL,
    "lastOperationTime" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.lastOperationTime = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON banking
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO "banking" (balance) VALUES (4000.0)
INSERT INTO "banking" (balance) VALUES (1000.0)
