SELECT author_id AS id FROM Views 
where author_id = viewer_id 
GROUP BY id
ORDER BY id

# second try
select distinct(author_id) as id
from Views
where author_id = viewer_id
order by id asc
