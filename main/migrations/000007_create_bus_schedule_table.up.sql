CREATE TABLE IF NOT EXISTS bus_schedule (
  id bigserial PRIMARY KEY,
  company_id BIGINT REFERENCES bus_companies(id),
  beginning_location_id BIGINT REFERENCES locations_of_stop(id),
  destination_location_id BIGINT REFERENCES locations_of_stop(id),
  bus_departure_time TIME,
  bus_arrival_time TIME
);