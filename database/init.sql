CREATE DATABASE habittrackerdb;

GRANT ALL PRIVILEGES ON DATABASE habittrackerdb TO lauti;

\c habittrackerdb

-- Auth Tables
CREATE TABLE auth (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  roles VARCHAR(255) NOT NULL,
  is_account_non_expired BOOLEAN,
  is_account_non_locked BOOLEAN,
  is_credentials_non_expired BOOLEAN,
  is_enabled BOOLEAN
);

-- Business logic tables
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  auth_id UUID NOT NULL,
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  image_url VARCHAR(1000) NOT NULL,
  FOREIGN KEY (auth_id) REFERENCES auth(id) ON DELETE CASCADE
);

CREATE TYPE day_of_week AS ENUM ('Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday');
CREATE TYPE time_of_day AS ENUM ('Morning', 'Afternoon', 'Evening');

CREATE TABLE activities (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  text VARCHAR(1000) NOT NULL,
  time_of_day time_of_day NOT NULL,
  day_of_week day_of_week NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TYPE difficulty AS ENUM ('Easy', 'Moderate', 'Hard');

CREATE TABLE routines (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  title VARCHAR(200) NOT NULL,
  difficulty difficulty NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE routines_activities (
  routine_id BIGINT NOT NULL,
  activity_id BIGINT NOT NULL ,
  PRIMARY KEY (routine_id, activity_id),
  FOREIGN KEY(routine_id) REFERENCES routines(id) ON DELETE CASCADE,
  FOREIGN KEY(activity_id) REFERENCES activities(id) ON DELETE CASCADE
);

CREATE TABLE users_activities (
  user_id BIGINT NOT NULL,
  activity_id BIGINT NOT NULL ,
  PRIMARY KEY (user_id, activity_id),
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY(activity_id) REFERENCES activities(id) ON DELETE CASCADE
);

CREATE TABLE users_routines (
  user_id BIGINT NOT NULL,
  routine_id BIGINT NOT NULL,
  PRIMARY KEY (user_id, routine_id),
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY(routine_id) REFERENCES activities(id) ON DELETE CASCADE
);

CREATE TABLE users_follows (
  follower BIGINT NOT NULL,
  followed BIGINT NOT NULL,
  PRIMARY KEY (follower, followed),
  FOREIGN KEY(follower) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY(followed) REFERENCES users(id) ON DELETE CASCADE
);
