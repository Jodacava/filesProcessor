DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions
(
    id               SERIAL NOT NULL PRIMARY KEY,
    month            VARCHAR NOT NULL,
    day              int NOT NULL,
    transaction      decimal   not null,
    email_to         VARCHAR   NOT NULL
);