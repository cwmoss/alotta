CREATE TABLE
    IF NOT EXISTS blogs (
        id INTEGER PRIMARY KEY,
        title text,
        content text,
        created_at datetime,
        modified_at datetime
    );