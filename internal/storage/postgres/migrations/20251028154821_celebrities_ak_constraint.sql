-- Drop index "ak_name_celebrity" from table: "celebrities"
DROP INDEX "ak_name_celebrity";
-- Create index "ak_celebrity" to table: "celebrities"
CREATE UNIQUE INDEX "ak_celebrity" ON "celebrities" ("name", "birth_date");
