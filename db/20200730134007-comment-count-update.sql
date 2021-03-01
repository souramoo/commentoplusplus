-- This trigger is called every time a comment is deleted, so the comment count for the page where the comment belong is updated
CREATE OR REPLACE FUNCTION commentsDeleteTriggerFunction() RETURNS TRIGGER AS $trigger$
BEGIN
  UPDATE pages
  SET commentCount = commentCount - 1
  WHERE domain = old.domain AND path = old.path;

  RETURN OLD;
END;
$trigger$ LANGUAGE plpgsql;

CREATE TRIGGER commentsDeleteTrigger AFTER DELETE ON comments
FOR EACH ROW EXECUTE PROCEDURE commentsDeleteTriggerFunction();

-- fix any broken comment counts
UPDATE pages SET commentCount = 0;
UPDATE pages
SET commentCount = subquery.commentCount
FROM (
  SELECT COUNT(commentHex) as commentCount, DOMAIN AS dmn, PATH AS pth
  FROM comments
  WHERE state = 'approved'
  GROUP BY (domain, path)
) as subquery
WHERE domain = subquery.dmn AND path=subquery.pth;