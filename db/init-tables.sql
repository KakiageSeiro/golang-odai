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
  password    varchar(256) not null
<<<<<<< HEAD
);

create table if not exists comments
(
  id      int unsigned not null primary key auto_increment,
  user_id    int not null,
  post_id    int not null,
  text    varchar(256) not null
);
=======
);
>>>>>>> 95a6fbd55e2ebbb8edba0eda21840f5764a6fe3a
