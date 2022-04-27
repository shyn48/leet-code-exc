DELETE p1 FROM Person p1 join Person p2 
where p1.email = p2.email and p1.id > p2.id

#alternate solution

DELETE FROM Person WHERE Id NOT IN 
(SELECT * FROM(
    SELECT MIN(Id) FROM Person GROUP BY Email) as p);
