CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     full_name VARCHAR(100)  NOT NULL,
    email VARCHAR(50) NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    balance MONEY  NOT NULL
    );

CREATE TABLE IF NOT EXISTS spaceships(
                                         id SERIAL PRIMARY KEY,
                                         name VARCHAR(100) NOT NULL,
    count_seats INT NOT NULL
    );

CREATE TABLE IF NOT EXISTS space_trips(
                                          id SERIAL PRIMARY KEY,
                                          from_planet VARCHAR(100) NOT NULL,
    to_planet VARCHAR(100) NOT NULL,
    price MONEY  NOT NULL,
    dispatch_date TIMESTAMPTZ NOT NULL,
    spaceship_id INT NOT NULL,
    UNIQUE(spaceship_id,dispatch_date),
    CONSTRAINT fk_spaceship FOREIGN KEY(spaceship_id) REFERENCES spaceships(id)
    );


CREATE TABLE IF NOT EXISTS client_trips(
                                           id SERIAL PRIMARY KEY,
                                           user_id INT NOT NULL,
                                           trip_id INT NOT NULL,
                                           purchase_date TIMESTAMPTZ NOT NULL,
                                           CONSTRAINT fk_travel FOREIGN KEY(trip_id) REFERENCES space_trips(id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)

    );