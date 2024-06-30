CREATE TABLE IF NOT EXISTS "Teacher" (
  "ID" uuid PRIMARY KEY,
  "FullName" varchar(100),
  "Phone" varchar(15),
  "Password" varchar(255),
  "Salary" int,
  "IeltsScore" int,
  "IeltsAttemptsCount" int,
  "SupportTeacherID" uuid,
  "BranchID" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "SupportTeacher" (
  "ID" uuid PRIMARY KEY,
  "FullName" varchar(100),
  "Phone" varchar(15),
  "Password" varchar(255),
  "Salary" int,
  "IeltsScore" int,
  "IeltsAttemptsCount" int,
  "BranchID" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "Administration" (
  "ID" uuid PRIMARY KEY,
  "FullName" varchar(100),
  "Phone" varchar(15),
  "Password" varchar(255),
  "Salary" int,
  "IeltsScore" int,
  "BranchID" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "Manager" (
  "ID" uuid PRIMARY KEY,
  "FullName" varchar(100),
  "Phone" varchar(15),
  "Password" varchar(255),
  "Salary" int,
  "BranchID" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "Branch" (
  "ID" uuid PRIMARY KEY,
  "Addres" varchar,
  "Location" polygon,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "Teacher" ADD COLUMN "LoginID" varchar UNIQUE;

ALTER TABLE "Manager" ADD COLUMN "LoginID" varchar UNIQUE;

ALTER TABLE "Administration" ADD COLUMN "LoginID" varchar UNIQUE;

ALTER TABLE "SupportTeacher" ADD COLUMN "LoginID" varchar UNIQUE;


ALTER TABLE "Teacher" ADD FOREIGN KEY ("SupportTeacherID") REFERENCES "SupportTeacher" ("ID");

ALTER TABLE "Teacher" ADD FOREIGN KEY ("BranchID") REFERENCES "Branch" ("ID");

ALTER TABLE "SupportTeacher" ADD FOREIGN KEY ("BranchID") REFERENCES "Branch" ("ID");

ALTER TABLE "Administration" ADD FOREIGN KEY ("BranchID") REFERENCES "Branch" ("ID");

ALTER TABLE "Manager" ADD FOREIGN KEY ("BranchID") REFERENCES "Branch" ("ID");

ALTER TABLE "Teacher" ADD COLUMN "Email" varchar(255);

ALTER TABLE "SupportTeacher" ADD COLUMN "Email" varchar(255);

ALTER TABLE "Administration" ADD COLUMN "Email" varchar(255);

ALTER TABLE "Manager" ADD COLUMN "Email" varchar(255);

CREATE EXTENSION IF NOT EXISTS postgis;
  
ALTER TABLE "Branch" 
ALTER COLUMN "Location" 
TYPE GEOMETRY(POINT, 4326) 
USING ST_SetSRID(ST_Centroid("Location"::geometry), 4326);
UPDATE "Branch"
SET "Location" = ST_GeomFromText("Location"::text, 4326);

ALTER TABLE "Branch"
ADD CONSTRAINT unique_branch_code UNIQUE ("Addres");