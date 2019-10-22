-- drop table if exists posts;
create table if not exists posts
(
  id      int unsigned not null primary key auto_increment,
  name    varchar(128) not null,
  text    varchar(256) not null
);

-- drop table if exists users;
create table if not exists users
(
  id      int unsigned not null primary key auto_increment,
  username    varchar(256) not null,
  authentication_id  varchar(256) not null unique
--   password    varchar(256) not null
);

create table if not exists comments
(
  id      int unsigned not null primary key auto_increment,
  user_id    int not null,
  post_id    int not null,
  text    varchar(256) not null
);
