SELECT
  Score,
  (SELECT count(*) FROM 
      (SELECT distinct Score s FROM Scores) tmp
        WHERE s >= Score) as 'rank'
FROM Scores
ORDER BY Score desc
