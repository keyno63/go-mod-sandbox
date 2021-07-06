DROP KEYSPACE IF EXISTS test;

CREATE KEYSPACE test WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor' : 1
};

USE test;

CREATE TABLE test_table (
    id text PRIMARY KEY,
    data text,
    update_time timestamp
) WITH default_time_to_live = 31536000;

INSERT INTO test.test_table (id, data, update_time)
VALUES('test_id', 'text', '2019-09-10T11:21:59.001+0000');
