-- SQLite
SELECT
    email, owner, 
    ifnull(count,0) as count,
    users.id as id 
from users
LEFT JOIN 
    (SELECT count(*) as count, owner from participations group by owner) 
    on owner = users.id
;

SELECT count(*), owner from participations group by owner;

SELECT length as length, owner as id from streaks ORDER BY lastweek DESC LIMIT 1;

SELECT
     users.id as id,  ifnull(len, 0) as length
from users
LEFT JOIN 
    (SELECT length as len, owner from streaks ORDER BY lastweek DESC LIMIT 1) 
    on owner = users.id
;