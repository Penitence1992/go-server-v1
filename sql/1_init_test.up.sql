-- 测试表

create table if not exists test_table
(
    id           varchar(255) not null
        constraint test_table_pkey primary key,
    name varchar (255) not null,
    created      timestamp,
    del      boolean,
    updated      timestamp
);