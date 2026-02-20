-- Migration to support multiple owners per domain

-- Create the domainOwners junction table
CREATE TABLE IF NOT EXISTS domainOwners (
  domain                   TEXT          NOT NULL                           ,
  ownerHex                 TEXT          NOT NULL                           ,
  addDate                  TIMESTAMP     NOT NULL                           ,
  PRIMARY KEY (domain, ownerHex)
);

-- Migrate existing domain-owner relationships to the new table
INSERT INTO domainOwners (domain, ownerHex, addDate)
SELECT domain, ownerHex, creationDate
FROM domains
WHERE ownerHex IS NOT NULL AND ownerHex != ''
ON CONFLICT (domain, ownerHex) DO NOTHING;

-- Remove the ownerHex column from domains table as it's now in domainOwners
ALTER TABLE domains DROP COLUMN IF EXISTS ownerHex;
