CREATE TABLE Loans (
                       LoanID SERIAL PRIMARY KEY,
                       BookID INT,
                       UserID INT,
                       LoanDate DATE,
                       ReturnDate DATE,
                       FOREIGN KEY (BookID) REFERENCES Books(BookID),
                       FOREIGN KEY (UserID) REFERENCES Users(UserID)
);