drop table if exists posts;
create table if not exists posts
(
  id      int unsigned not null primary key auto_increment,
  user_id    varchar(128) not null,
  text    varchar(256) not null
);

drop table if exists users;
create table if not exists users
(
  id      int unsigned not null primary key auto_increment,
  session_id    varchar(256) not null,
  username    varchar(256) not null
);

drop table if exists comments;
create table if not exists comments
(
  id      int unsigned not null primary key auto_increment,
  user_id    varchar(256) not null,
  post_id    int not null,
  text    varchar(256) not null
);
