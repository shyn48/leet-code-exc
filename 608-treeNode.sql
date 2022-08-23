# Write your MySQL query statement below

select id, 'Root' as type from Tree
where p_id is null

union

select id, 'Leaf' as type from Tree
where
id not in (
    select distinct p_id from Tree where
    p_id is not null
) and p_id is not null

union 

select id, 'Inner' as type From Tree
where 
id in (
    select distinct p_id from Tree where
    p_id is not null
) and p_id is not null 

#alternative 

SELECT
    id AS `Id`,
    CASE
        WHEN tree.id = (SELECT atree.id FROM tree atree WHERE atree.p_id IS NULL)
          THEN 'Root'
        WHEN tree.id IN (SELECT atree.p_id FROM tree atree)
          THEN 'Inner'
        ELSE 'Leaf'
    END AS Type
FROM
    tree
ORDER BY `Id`
;
