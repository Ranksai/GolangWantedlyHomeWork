CREATE TABLE wantedlyusers (
 id serial,
 name varchar(255) not null,
 email varchar(255) not null unique,
 created_at timestamp default CURRENT_TIMESTAMP,
 updated_at timestamp default current_timestamp,
 primary key(id)
);
create function set_update_time() returns opaque as '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' language 'plpgsql';

create trigger update_tri before update on wantedlyusers  for each row
  execute procedure set_update_time();
