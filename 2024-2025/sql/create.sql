SET SEARCH_PATH TO course;

-- Table structure for `companies`
DROP TABLE IF EXISTS companies;
CREATE TABLE companies
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(60)
);

-- Insert data into `companies`
INSERT INTO companies (id, name)
VALUES (1, 'ABC Ltd.'),
       (2, 'Sofia University'),
       (3, 'Best Widgets Ltd.'),
       (4, 'Software AD');

-- Table structure for `projects`
DROP TABLE IF EXISTS projects;
CREATE TABLE projects
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(60)      NOT NULL UNIQUE,
    description VARCHAR(1024),
    budget      DOUBLE PRECISION NOT NULL,
    finished    BOOLEAN          NOT NULL,
    start_date  DATE,
    company_id  BIGINT REFERENCES companies (id)
);

-- Insert data into `projects`
INSERT INTO projects (id, name, description, budget, finished, start_date, company_id)
VALUES (1, 'Build CI/CD Server', 'Build custom continuous integration server for our projects ...', 70000, FALSE,
        '2021-01-06', 4),
       (2, 'Create Furniture Web Site', 'Build web site for our client selling furniture ...', 20000, FALSE,
        '2021-01-06', 1),
       (3, 'Update SUSI with eLearning Functionality', 'Add eLearning functionality to SUSI ...', 50000, FALSE,
        '2021-01-06', 3),
       (4, 'Build IoT Control Access System', 'Build custom IoT system controlling access to FMI building ...', 70000,
        FALSE, '2021-01-06', 3),
       (5, 'Create Thymeleaf App', 'Create Thymeleaf App with Spring Boot and Spring Data', 1200, FALSE, '2021-01-12',
        2),
       (6, 'Test Project', 'Angular challenge for everybody', 1, FALSE, '2021-01-12', 2),
       (7, 'Learn Golang DBs', 'Do homework to learn Golang database management', 50, FALSE, '2021-01-13', 2);

-- Table structure for `users`
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(20),
    last_name  VARCHAR(20),
    email      VARCHAR(255) UNIQUE,
    username   VARCHAR(30) UNIQUE,
    password   VARCHAR(255),
    active     BOOLEAN NOT NULL,
    created    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert data into `users`
INSERT INTO users (id, first_name, last_name, email, username, password, active, created, modified)
VALUES


-- Table structure for `projects_users`
DROP TABLE IF EXISTS projects_users;
CREATE TABLE projects_users
(
    project_id BIGINT REFERENCES projects (id),
    user_id    BIGINT REFERENCES users (id),
    PRIMARY KEY (project_id, user_id)
);

-- Insert data into `projects_users`
INSERT INTO projects_users (project_id, user_id)
VALUES (1, 1),
       (2, 1),
       (3, 1),
       (1, 2),
       (3, 2),
       (1, 3),
       (2, 3);

-- Table structure for `user_roles`
DROP TABLE IF EXISTS user_roles;
CREATE TABLE user_roles
(
    user_id BIGINT REFERENCES users (id),
    role    VARCHAR(255)
);

-- Insert data into `user_roles`
INSERT INTO user_roles (user_id, role)
VALUES (1, 'ADMIN'),
       (2, 'EMPLOYEE'),
       (3, 'MANAGER'),
       (3, 'ADMIN'),
       (3, 'EMPLOYEE');

(1, 'Default', 'Admin', 'admin@gmail.com', 'admin', '{bcrypt}$2a$10$wtZZE9untaIrJgR2P9klsuBVdVVxd0QM1Z..R1aE0YsucS.IkXIu.', FALSE, '2021-01-06 22:22:40.554', '2021-01-06 22:22:40.554')
,
                                                                                                        (2, 'Ivan', 'Petrov', 'ivan@gmail.com', 'ivan', '{bcrypt}$2a$10$pm2OKA/ESO3rpGvI0YZOVOqrTvl1HUhyAAOi.ztUvK/K7xca1aKMy', FALSE, '2021-01-06 22:22:40.554', '2021-01-06 22:22:40.554'),
                                                                                                        (3, 'Veronika', 'Dimitrova', 'vera@gmail.com', 'veronika', '{bcrypt}$2a$10$ls9amuWqw39yXgX4s20DDecgCOEZXx1PPCPuINizF1rzTmG0vzPLG', FALSE, '2021-01-06 22:22:40.554', '2021-01-06 22:22:40.554');
