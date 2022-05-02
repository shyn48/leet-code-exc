select player_id, min(event_date) first_login
from Activity
group by player_id;

# also learned about dence_rank 

select player_id, first_login
from (
    select 
    player_id, 
    event_date first_login,
    dense_rank() over(
        partition by player_id
        order by event_date
    ) as row_rank
    from Activity
) as subquery
where row_rank = 1;
