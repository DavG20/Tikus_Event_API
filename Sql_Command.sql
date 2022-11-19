Cretae DATABASE tikus_event;

Create Table user_tikus_event (
    user_id Serial PRIMARY KEY,
    user_name Text  not NULL UNIQUE ,
    email text NOT NULL UNIQUE ,
    password text not NULL,
    created_on Time not NULL DEFAULT CURRENT_TIMESTAMP,
    profile_url text,
    admin Boolean not null DEFAULT False

);
insert into user_tikus_event(user_name,email,password,created_on,profile_url)  VALUES ('dav','dawit@gmail.com','dav1234',CURRENT_TIMESTAMP,'image/profile.jpg');