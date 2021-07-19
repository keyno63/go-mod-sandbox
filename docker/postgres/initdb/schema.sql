CREATE TABLE IF NOT EXISTS sandbox_table (
    id integer NOT NULL,
    name varchar NOT NULL,
    category varchar NOT NULL
);

INSERT INTO sandbox_table VALUES (1, 'taro', 'category1');
INSERT INTO sandbox_table VALUES (2, 'jiro', 'category1');
INSERT INTO sandbox_table VALUES (3, 'saburo', 'category2');

CREATE TABLE IF NOT EXISTS sandbox_json_table (
     id integer NOT NULL,
     updated_at json NOT NULL
);

INSERT INTO sandbox_json_table VALUES (1, '{"Nanos": 5353, "Seconds": 234}');

