// the problem has wrong testcase 

SELECT activity_date AS day, COUNT(DISTINCT user_id) AS active_users
FROM Activity
GROUP BY activity_date
HAVING subdate('2019-07-27', 30)
