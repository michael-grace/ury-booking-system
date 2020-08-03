create schema bookings;

comment on schema bookings is 'For the URY booking system.';

create table bookings.request_levels(
    request_level int unique,
    description varchar
);

create table bookings.resources(
    resource SERIAL,
    name varchar,
    PRIMARY KEY (resource)
);

create table bookings.unique_resources(
    unique_id SERIAL,
    resource int,
    name varchar,
    primary key (unique_id),
    foreign key (resource) references bookings.resources(resource)
);

create table bookings.public_contacts(
    public_id SERIAL,
    contact varchar,
    primary key (public_id)
);

create table bookings.bookings(
    booking_id SERIAL,
    member_id int,
    request_level int,
    resource int,
    preference int,
    given_resource int,
    timeslot_id int,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    public_id int,
    application_datetime timestamp with time zone,
    PRIMARY KEY (booking_id),
    foreign key (member_id) references public.member(memberid),
    foreign key (request_level) references bookings.request_levels(request_level),
    foreign key (resource) references bookings.resources(resource),
    foreign key (preference) references bookings.unique_resources(unique_id),
    foreign key (given_resource) references bookings.unique_resources(unique_id),
    --foreign key (timeslot_id) references schedule.show_season_timeslot(show_season_timeslot_id), REMOVED SO TIMESLOT CAN BE SCHEDULED SIMULTANEOUSLY
    foreign key (public_id) references bookings.public_contacts(public_id)
);

INSERT INTO bookings.public_contacts (public_contacts.public_id, public_contacts.contact) VALUES (0, "Not Public")
INSERT INTO bookings.unique_resources (unique_resources.unique_id, unique_resources.description) VALUES (0, "No Preference")