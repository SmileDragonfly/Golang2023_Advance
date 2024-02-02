USE gormdb;
CREATE PARTITION FUNCTION users_pf (int) AS RANGE LEFT FOR VALUES (1000, 2000, 3000, 4000, 5000);
CREATE PARTITION SCHEME users_ps AS PARTITION users_pf ALL TO ([PRIMARY]);
--ALTER TABLE users ON users_ps (id);
