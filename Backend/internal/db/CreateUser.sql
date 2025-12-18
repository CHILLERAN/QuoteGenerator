CREATE USER 'quotereader'@'localhost';
GRANT SELECT ON randomquotesdb.* TO 'quotereader'@'localhost';
ALTER USER 'quotereader'@'localhost' IDENTIFIED BY 'password';