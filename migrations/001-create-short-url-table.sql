create table if not exists shorturl (
                                        id bigserial primary key ,
                                        url text not null,
                                        short_code text unique,
                                        created_by text,
                                        created_at timestamp not null
)