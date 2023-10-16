select count, id from
  (
  SELECT COUNT(*) as count, region as id FROM users GROUP BY region
  UNION ALL
  SELECT COUNT(*) as count, "usercount" as id FROM users
  )
;