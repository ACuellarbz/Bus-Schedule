CREATE TABLE IF NOT EXISTS bus_schedule (
  id bigserial PRIMARY KEY,
  route_name BIGINT REFERENCES routes(id) NOT NULL,
  total_cost BIGINT NOT NULL,
  number_of_tickets_available BIGINT NOT NULL
);