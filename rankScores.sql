SELECT
  Score,
  (SELECT count(*) FROM 
      (SELECT distinct Score s FROM Scores) tmp
        WHERE s >= Score) as 'rank'
FROM Scores
ORDER BY Score desc

# second try
Select 
    s.Score,
    (select count(distinct score) from Scores s2 where s2.Score >= s.Score) as 'Rank'
    from Scores s
Order by s.Score desc
