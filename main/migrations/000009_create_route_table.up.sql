CREATE TABLE IF NOT EXISTS routes (
  id bigserial PRIMARY KEY,
  route_name BIGINT REFERENCES bus_schedule(id),
  number_of_miles BIGINT,
  total_cost BIGINT,
  number_of_tickets_available BIGINT 
);