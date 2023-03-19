CREATE TABLE users 
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    salt varchar(24) not null
);

CREATE TABLE projects 
(
    id serial primary key,
    name varchar(255) not null
);

CREATE TABLE projects_users
(
    id serial primary key,
    project_id bigint not null,
    user_id bigint not null,
    access varchar(10) not null,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE cars 
(
    id serial primary key,
    project_id bigint not null,
    name varchar(255) not null,
    load_capacity real not null,
    width real not null,
    height real not null,
    length real not null,
    fuel_consumption real not null,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE clients 
(
    id serial primary key,
    project_id bigint not null,
    name varchar(255) not null,
    address varchar(255) not null,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE cargos 
(
    id serial primary key,
    project_id bigint not null,
    name varchar(255) not null,
    width real not null,
    height real not null,
    length real not null,
    weight real not null,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

CREATE TABLE orders 
(
    id serial primary key,
    client_id bigint not null,
    cargo_id bigint not null,
    count int not null,
    FOREIGN KEY (client_id) REFERENCES clients (id) ON DELETE CASCADE,
    FOREIGN KEY (cargo_id) REFERENCES cargos (id) ON DELETE CASCADE
);