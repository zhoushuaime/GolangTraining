create database `test` CHARACTER set utf8mb4;
create table test(
                     id BIGINT(20) not null auto_increment,
                     a int not null,
                     b int not null,
                     PRIMARY key(id)
)engine=INNODB;


INSERT INTO test(a,b) VALUES(0,2);
INSERT INTO test(a,b) VALUES(0,4);
INSERT INTO test(a,b) VALUES(0,7);
INSERT INTO test(a,b) VALUES(0,8);