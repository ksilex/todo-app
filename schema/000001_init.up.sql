create table users (
  id serial primary key,
  name varchar(255) not null,
  username varchar(255) not null unique,
  password_hash varchar(255) not null
);

create table todo_lists (
  id serial primary key,
  name varchar(255) not null,
  description varchar(255)
);

create table users_lists (
  id serial primary key,
  user_id int references users (id) on delete cascade not null,
  todo_id int references todo_lists (id) on delete cascade not null
);

create table todo_items (
  id serial primary key,
  name varchar(255) not null,
  description varchar(255),
  done boolean not null default false
);

create table items_lists (
  id serial primary key,
  item_id int references todo_items (id) on delete cascade not null,
  todo_id int references todo_lists (id) on delete cascade not null
);
