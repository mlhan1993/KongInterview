CREATE DATABASE IF NOT EXISTS kong;

CREATE TABLE IF NOT EXISTS kong.services (
    id INT NOT NULL AUTO_INCREMENT, 
    name VARCHAR(255) NOT NULL UNIQUE, 
    description VARCHAR(255), 
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS kong.versions(
    id INT NOT NULL AUTO_INCREMENT,
    tag VARCHAR(8) NOT NULL,
    serviceID INT NOT NULL,
    dateCreated DATETIME,
    
    PRIMARY KEY (id),
    UNIQUE KEY service_tag (serviceID, tag),
    FOREIGN KEY (serviceID) REFERENCES services(id)
);


INSERT IGNORE INTO kong.services (name, description) VALUES 
("A","Service A"),
("B","Service B"),
("C","Service C"),
("D","Service D"),
("E","Service E"),
("F","Service F"),
("G","Service G"),
("H","Service H"),
("I","Service I"),
("J","Service J"),
("K","Service K"),
("L","Service L"),
("M","Service M"),
("N","Service N"),
("O","Service O"),
("P","Service P"),
("Q","Service Q"),
("R","Service R"),
("S","Service S"),
("T","Service T"),
("U","Service U"),
("V","Service V"),
("W","Service W"),
("X","Service X"),
("Y","Service Y"),
("Z","Service Z"),
("AA","Service AA"),
("AB","Service AB"),
("AC","Service AC"),
("AD","Service AD");


INSERT IGNORE INTO kong.versions (serviceID, tag, dateCreated)
VALUES
  (1, "v1", "2023-01-01"),
  (1, "v2", "2023-01-02"),
  (1, "v3", "2023-01-03"),
  (1, "v4", "2023-01-04"),
  (1, "v5", "2023-01-05"),
  (2, "v1", "2023-01-01"),
  (2, "v2", "2023-01-02"),
  (2, "v3", "2023-01-03"),
  (2, "v4", "2023-01-04"),
  (3, "v1", "2023-01-01"),
  (3, "v2", "2023-01-02"),
  (4, "v1", "2023-01-01"),
  (4, "v2", "2023-01-02"),
  (4, "v3", "2023-01-03"),
  (4, "v4", "2023-01-04"),
  (4, "v5", "2023-01-05"),
  (5, "v1", "2023-01-01"),
  (5, "v2", "2023-01-02"),
  (5, "v3", "2023-01-03"),
  (5, "v4", "2023-01-04"),
  (5, "v5", "2023-01-05"),
  (5, "v6", "2023-01-06"),
  (5, "v7", "2023-01-07"),
  (6, "v1", "2023-01-01"),
  (6, "v2", "2023-01-02"),
  (7, "v2", "2023-01-02"),
  (8, "v2", "2023-01-02"),
  (9, "v2", "2023-01-02"),
  (10, "v2", "2023-01-02"),
  (11, "v2", "2023-01-02"),
  (12, "v2", "2023-01-02"),
  (13, "v2", "2023-01-02"),
  (14, "v2", "2023-01-02"),
  (15, "v2", "2023-01-02"),
  (16, "v2", "2023-01-02"),
  (17, "v2", "2023-01-02"),
  (18, "v2", "2023-01-02"),
  (19, "v2", "2023-01-02"),
  (20, "v2", "2023-01-02"),
  (21, "v2", "2023-01-02"),
  (22, "v2", "2023-01-02"),
  (23, "v2", "2023-01-02"),
  (24, "v2", "2023-01-02"),
  (25, "v2", "2023-01-02"),
  (26, "v2", "2023-01-02"),
  (27, "v2", "2023-01-02"),
  (28, "v2", "2023-01-02"),
  (29, "v2", "2023-01-02"),
  (30, "v2", "2023-01-02");
