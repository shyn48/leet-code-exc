# Write your MySQL query statement below

select name, sum(amount) as balance from 
Users u join Transactions t on u.account = t.account
group by u.account
having sum(amount) > 10000


#second attempt
select name, sum(amount) as balance
from Users u
     Join Transactions t
     on u.account = t.account
group by u.account 
having balance > 10000
