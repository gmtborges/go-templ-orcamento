-- +goose Up
-- +goose StatementBegin
CREATE TABLE employers (
  id SERIAL PRIMARY KEY,
  type VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CHECK (type IN ('company', 'auto_store'))
);

CREATE TABLE companies (
  id SERIAL PRIMARY KEY,
  employer_id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255),
  contact_number INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (employer_id) REFERENCES employers (id) ON DELETE NO ACTION
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(50) NOT NULL,
  employer_id INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (employer_id) REFERENCES employers (id) ON DELETE NO ACTION
);

CREATE TABLE auto_stores (
  id SERIAL PRIMARY KEY,
  employer_id INTEGER NOT NULL,
  name VARCHAR(255) NOT NULL,
  location VARCHAR(255),
  contact_number INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (employer_id) REFERENCES employers (id) ON DELETE NO ACTION
);

CREATE TABLE auto_categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE auto_stores_categories (
  auto_category_id INTEGER REFERENCES auto_categories (id),
  auto_store_id INTEGER REFERENCES auto_stores (id),
  PRIMARY KEY (auto_category_id, auto_store_id)
);

CREATE TABLE biddings (
  id SERIAL PRIMARY KEY,
  company_id INTEGER NOT NULL,
  title VARCHAR(200) NOT NULL,
  description TEXT,
  start_date TIMESTAMP,
  end_date TIMESTAMP,
  category_id INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE NO ACTION,
  FOREIGN KEY (category_id) REFERENCES auto_categories (id) ON DELETE NO ACTION
);

CREATE TABLE auto_offers (
  id SERIAL PRIMARY KEY,
  bidding_id INTEGER NOT NULL,
  auto_store_id INTEGER NOT NULL,
  offer_details TEXT,
  offer_date TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (bidding_id) REFERENCES biddings (id) ON DELETE NO ACTION,
  FOREIGN KEY (auto_store_id) REFERENCES auto_stores (id) ON DELETE NO ACTION
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE auto_offers;

DROP TABLE auto_stores_categories;

DROP TABLE biddings;

DROP TABLE auto_stores;

DROP TABLE companies;

DROP TABLE users;

DROP TABLE auto_categories;

DROP TABLE employers;

-- +goose StatementEnd
