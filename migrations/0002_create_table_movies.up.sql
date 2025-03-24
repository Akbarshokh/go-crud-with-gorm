CREATE TABLE movies (
                        id SERIAL PRIMARY KEY,
                        title TEXT NOT NULL,
                        director TEXT NOT NULL,
                        year INTEGER NOT NULL,
                        plot TEXT,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        deleted_at TIMESTAMP
);