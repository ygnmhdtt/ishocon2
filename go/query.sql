-- userがvoteしたcount
alter table users add column voted int(8) default 0 not null;

-- votesテーブルはcandidate_idとcountのみ
drop table votes;
create table votes (candidate_id int(11) primary key, vote_count int(8) default 0 not null) ;

-- keywordテーブル
create table keyword (candidate_id int(11), keyword varchar(255) default "" not null, count int(8) default 0 not null);
