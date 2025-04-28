create database if not exists Stacker;
use Stacker;
create table if not exists Stacks (
  id int not null auto_increment,
  title varchar(50),
  body varchar(1000),
  creationDate timestamp default current_timestamp,
  updateDate timestamp default current_timestamp on update CURRENT_TIMESTAMP,

  primary key (id)
)

