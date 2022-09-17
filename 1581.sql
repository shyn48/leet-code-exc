SELECT customer_id, COUNT(*) as count_no_trans
FROM Visits
WHERE visit_id NOT IN (SELECT DISTINCT visit_id FROM Transactions)
GROUP BY customer_id;

# or

SELECT customer_id, COUNT(Visits.visit_id) AS count_no_trans
FROM
	visitsLEFT JOIN Transactions
	ON Visits.visit_id = Transactions.visit_id
WHERE 
	Transactions.visit_id IS NULL
GROUP BY customer_id


#second try

select customer_id, count(customer_id) as count_no_trans
from Visits v left join Transactions t on v.visit_id = t.visit_id
where transaction_id is null
group by customer_id
