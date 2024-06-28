CREATE TABLE Books (
                       BookID SERIAL PRIMARY KEY,
                       Title VARCHAR(255),
                       Author VARCHAR(255),
                       PublicationYear INT,
                       Genre VARCHAR(100),
                       Status VARCHAR(10)
)
