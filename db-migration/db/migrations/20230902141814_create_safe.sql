-- migrate:up
CREATE TABLE safe (
    id integer primary key autoincrement, 
    user varchar(255),
    key varchar(255),
    secret blob
);


-- migrate:down
DROP TABLE IF EXISTS safe;
