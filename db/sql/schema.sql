CREATE TABLE IF NOT EXISTS todos (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
		folder text NOT NULL,
    description TEXT NOT NULL,
    done        INTEGER NOT NULL DEFAULT 0
);
