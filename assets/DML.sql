-- Insert data into mst_employee
INSERT INTO mst_employee (name, email, password, division, position, role, contact)
VALUES
    ('Budi', 'budi@mail.com', '12345', 'IT', 'Developer', 'EMPLOYEE', '1234567890'),
    ('Agung', 'agung@mail.com', '12345', 'HR', 'Manager', 'ADMIN', '9876543210'),
    ('Boby', 'boby@mail.com', '12345', 'Finance', 'Accountant', 'GA', '5678901234');

-- Insert data into mst_room
INSERT INTO mst_room (code_room, room_type, capacity, facilities)
VALUES
    ('R001', 'Meeting Room', 25, 'Projector, Whiteboard, Audio System, Video Conferencing'),
    ('R002', 'Conference Room', 20, 'Audio System, Video Conferencing, Large Screen Display'),
    ('R003', 'Event Hall', 50, 'Stage, Catering Services, Audio System'),
    ('R004', 'Boardroom', 30, 'Presentation Screen, Teleconferencing, Whiteboard');

-- Insert data into mst_facilities (continued)
INSERT INTO mst_facilities (code_name, facilities_type)
VALUES
    ('PRJ3', 'projector'),
    ('SCR1', 'screen'),
    ('SCR2', 'screen'),
    ('AUD1', 'audio system'),
    ('AUD2', 'audio system'),
    ('CAT1', 'catering services'),
    ('STG1', 'stage'),
    ('WHT1', 'whiteboard'),
    ('WHT2', 'whiteboard'),
    ('VID1', 'video conferencing'),
    ('VID2', 'video conferencing');

-- Insert data into tx_room_reservation
INSERT INTO tx_room_reservation (employee_id, room_id, start_date, end_date, notes, approval_status, approval_note)
VALUES
    ((SELECT id FROM mst_employee WHERE email = 'budi@mail.com'),
     (SELECT id FROM mst_room WHERE code_room = 'R001'),
     '2024-01-25 09:00:00', '2024-01-27 11:00:00', 'Team Meeting', 'ACCEPT', 'Department Briefing'),

    ((SELECT id FROM mst_employee WHERE email = 'agung@mail.com'),
     (SELECT id FROM mst_room WHERE code_room = 'R002'),
     '2024-01-26 14:00:00', '2024-01-26 16:00:00', 'Interview', 'PENDING', 'Training Session'),

    ((SELECT id FROM mst_employee WHERE email = 'boby@mail.com'),
    (SELECT id FROM mst_room WHERE code_room = 'R003'),
    '2024-01-27 10:00:00', '2024-01-27 12:00:00', 'Training Session', 'DECLINE', 'Board Meeting');

-- Insert data into tx_additional
INSERT INTO tx_additional (reservation_id, facilities_id)
VALUES
    ((SELECT id FROM tx_room_reservation WHERE approval_note = 'Department Briefing'),
     (SELECT id FROM mst_facilities WHERE code_name = 'PRJ3')),

    ((SELECT id FROM tx_room_reservation WHERE approval_note = 'Training Session'),
     (SELECT id FROM mst_facilities WHERE code_name = 'SCR1')),

    ((SELECT id FROM tx_room_reservation WHERE approval_note = 'Board Meeting'),
     (SELECT id FROM mst_facilities WHERE code_name = 'SCR2'));
