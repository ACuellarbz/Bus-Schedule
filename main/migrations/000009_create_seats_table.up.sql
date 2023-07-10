CREATE TABLE IF NOT EXISTS seats (
  id bigserial PRIMARY KEY,
  seat_name varchar(255) NOT NULL,
  type_of_seat varchar(255) NOT NULL  
);
