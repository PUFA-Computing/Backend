CREATE TABLE IF NOT EXISTS events (
                                      id SERIAL PRIMARY KEY,
                                      title TEXT NOT NULL,
                                      description TEXT NOT NULL,
                                      start_date DATE,
                                      end_date DATE,
                                      user_id uuid NOT NULL,
                                      status TEXT NOT NULL,
                                      link TEXT NOT NULL,
                                      thumbnail TEXT,
                                      created_at TIMESTAMP DEFAULT NOW(),
                                      updated_at TIMESTAMP DEFAULT NOW(),
                                      FOREIGN KEY (user_id) REFERENCES users (id)
);