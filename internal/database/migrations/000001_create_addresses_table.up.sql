CREATE TABLE addresses (
  id SERIAL NOT NULL PRIMARY KEY,
  country VARCHAR(255) NOT NULL,
  uf CHAR(2) NOT NULL,
  city VARCHAR(255) NOT NULL,
  street VARCHAR(255) NOT NULL,
  zipcode CHAR(8) NOT NULL,
  district VARCHAR(255),
  complement VARCHAR(255)
);