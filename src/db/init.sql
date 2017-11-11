CREATE TABLE wantedlyusers (
 id int primary key AUTO_INCREMENT,
 name varchar(255) not null,
 email varchar(255) not null,
 created_at timestamp not null default current_timestamp,
 updated_at timestamp not null default current_timestamp
);

create function set_update_time() returns opaque as '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' language 'plpgsql';

create trigger update_tri before update on posts for each row
  execute procedure set_update_time();