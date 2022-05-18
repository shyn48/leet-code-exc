# Write your MySQL query statement below
select
    (
        case
        when id % 2 != 0 and countId != id then id + 1
        when id % 2 != 0 and countId = id then id
        else id - 1
        end
    ) as id,
    student
    from seat,
    (select count(*) as countId from seat) as seat_count 
    order by id asc
