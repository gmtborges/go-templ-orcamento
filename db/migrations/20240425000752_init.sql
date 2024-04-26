-- +goose Up
-- +goose StatementBegin
-- Create 'companies' table
-- Create 'companies' table
CREATE TABLE companies (
  company_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  address TEXT,
  contact_info TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create 'auto_parts_categories' table
CREATE TABLE auto_parts_categories (
  category_id INTEGER PRIMARY KEY AUTOINCREMENT,
  category_name TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create 'biddings' table
CREATE TABLE biddings (
  bidding_id INTEGER PRIMARY KEY AUTOINCREMENT,
  company_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  start_date DATE,
  end_date DATE,
  category_id INTEGER NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies (company_id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES auto_parts_categories (category_id) ON DELETE CASCADE
);

-- Create 'auto_stores' table
CREATE TABLE auto_stores (
  store_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  location TEXT,
  contact_info TEXT,
  category_id INTEGER NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES auto_parts_categories (category_id) ON DELETE CASCADE
);

-- Create 'auto_offers' table
CREATE TABLE auto_offers (
  offer_id INTEGER PRIMARY KEY AUTOINCREMENT,
  bidding_id INTEGER NOT NULL,
  store_id INTEGER NOT NULL,
  offer_details TEXT,
  offer_date DATE,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (bidding_id) REFERENCES biddings (bidding_id) ON DELETE CASCADE,
  FOREIGN KEY (store_id) REFERENCES auto_stores (store_id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE companies;

DROP TABLE auto_parts_categories;

DROP TABLE biddings;

DROP TABLE auto_stores;

DROP TABLE auto_offers;

-- +goose StatementEnd
