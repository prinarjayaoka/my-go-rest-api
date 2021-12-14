CREATE TABLE threads
(
  id  varchar(64) not null primary key,
  title varchar(512)  not null,
  `description` varchar(2048) null
);

CREATE TABLE posts
(
    id varchar(64) not null primary key,
	thread_id varchar(64) not null,
	title varchar(512) not null,
	content varchar(2048) not null,
	votes int default 0 null
);

CREATE TABLE comments
(
    id int auto_increment primary key,
	post_id varchar(64) not null,
	content varchar(2048) not null,
	votes int default 0 null
);

