# Lucky Draw System

The Lucky Draw System is a web-based system that allows customers to participate in lucky draws and win prizes. It is written in Golang using the Gin framework, and the sample customers, prizes and mobile database is implemented using PostgreSQL.

## Running the System

To run the Lucky Draw System, follow these steps:

1. Navigate to the stack directory.
2. Execute the following command in your terminal:

## Command

``` sh
cd stack
docker-compose up
```

## Inspecting the Database

To inspect the database, you can use the following commands:

1. Open a shell in the "luckydraw_db" Docker container:

    ``` sh
    docker exec -it luckydraw_db sh
    ```

2. Enter the PostgreSQL database named "luckydraw":

    ``` sh
    psql -h localhost -p 5432 -U postgres -d luckydraw
    ```

3. Now, you can execute SQL queries to inspect the database. Here are some examples:

   - Inspect the CUSTOMERS table:

   ``` sh
   SELECT * FROM CUSTOMERS;
   ```

   This query will display the id and drawed columns for all rows in the CUSTOMERS table. The drawed column indicates whether a customer has been drawn or not.

   Original data:

   ``` sh
   # id | drawed 
   #----+--------
   #  3 | f
   #  4 | f
   #  1 | f
   #  2 | f
    ```

   - Inspect the CUSTOMERS table:

   ``` sh
   SELECT * FROM PRIZES;
    ```

   This query will display the category, probability, totalquota, dailyquota, and id columns for all rows in the PRIZES table.

   Original data:

   ``` sh
   #         category         | probability | totalquota | dailyquota | id 
   # -------------------------+-------------+------------+------------+----
   #  Buy 1 Get 1 Free Coupon |         0.8 |       9999 |       9999 |  4
   #  No prize                |       0.175 |       9999 |       9999 |  3
   #  $5 Cash Coupon          |       0.005 |        500 |        100 |  2
   #  $2 Cash Coupon          |        0.02 |       5000 |        500 |  1
   # (4 rows)
   ```

   - Inspect the MOBILE table:

   ``` sh
   SELECT * FROM MOBILE;
   ```

   This query will display the id, customerid, and mobile columns for all rows in the MOBILE table. It shows the mobile phone numbers associated with each customer.

   Original data:

   ``` sh
   #  id | customerid | mobile 
   # ----+------------+--------
   # (0 rows)
   ```

   - To adjust the probabilities for each prize category, you can use an UPDATE statement. For example, to update the probability of winning the "$2 Cash Coupon" to 0.5:

   ```sh
   UPDATE PRIZES SET probability = 0.5 WHERE id= 1;
   ```
