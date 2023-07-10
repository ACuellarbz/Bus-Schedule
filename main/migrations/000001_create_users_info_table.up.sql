CREATE TABLE IF NOT EXISTS users_info (
  id bigserial PRIMARY KEY,
  fname varchar(255) NOT NULL,
  lname varchar(255) NOT NULL,
  email citext UNIQUE NOT NULL,
  addres varchar(255) NOT NULL, --name change probably
  phone_number varchar(255) NOT NULL,
  passwrd bytea NOT NULL, --name change probably
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()

);