insert into cities (name) values
    ('Moscow'), ('Sochi'), ('Almaty'), ('Minsk'),
    ('Saint Petersburg'), ('Varna'), ('Novosibirsk'), ('Kazan'),
    ('Astana'), ('Tver');

insert into hotels (name, city_id, address, stars) values
  ('Grand Plaza', (select id from cities where name = 'Moscow'), 'ул. Тверская, 12, Москва', 5),
  ('Seaside Resort', (select id from cities where name = 'Sochi'), 'наб. Черноморская, 45, Сочи', 4),
  ('Mountain View Lodge', (select id from cities where name = 'Almaty'), 'ул. Горная, 8, Алматы', 3),
  ('Central Park Hotel', (select id from cities where name = 'Moscow'), 'пр. Свободы, 23, Минск', 4),
  ('Royal Palace', (select id from cities where name = 'Saint Petersburg'), 'пр. Невский, 56, Санкт-Петербург', 5),
  ('Sunset Beach Inn', (select id from cities where name = 'Varna'), 'пляж Золотые Пески, 77, Варна', 3),
  ('Business Center', (select id from cities where name = 'Novosibirsk'), 'ул. Вокзальная, 15, Новосибирск', 4),
  ('Family Cottage', (select id from cities where name = 'Kazan'), 'ул. Загородная, 32, Казань', 2),
  ('Luxury Tower', (select id from cities where name = 'Astana'), 'пр. Набережный, 88, Астана', 5),
  ('Economy Hostel', (select id from cities where name = 'Tver'),'ул. Молодежная, 10, Тверь', 1),
  ('Smirnovs Palace', (select id from cities where name = 'Saint Petersburg'), 'пр.Большой, 15, Санкт-Петербург', 3);