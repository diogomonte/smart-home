CREATE TABLE IF NOT EXISTS devices (
    id int auto_increment primary key not null,
    device_id character varying (50),
    device_type character varying (25)
);