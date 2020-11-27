CREATE KEYSPACE blog
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};


CREATE TABLE blog.posts (
  id uuid,
	user_id uuid,
	content text,
	PRIMARY KEY (id)
);

CREATE TABLE blog.users (
  id uuid,
	firstname text,
	lastname text,
	age int,
	email text,
	city text,
	PRIMARY KEY (id)
);
