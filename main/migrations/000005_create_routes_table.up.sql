CREATE TABLE IF NOT EXISTS routes (
  id bigserial PRIMARY KEY,
  beginning_location_id BIGINT REFERENCES locations_of_stop(id) NOT NULL,
  destination_location_id BIGINT REFERENCES locations_of_stop(id) NOT NULL,
  type_of_trip varchar(255) NOT NULL,
  bus_departure_time TIME NOT NULL,
  bus_arrival_time TIME NOT NULL
);