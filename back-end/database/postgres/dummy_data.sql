-- This script isn't working, and I'm not sure why
set session my.pokedex_number = '20000';
set session my.name = 'Dummy Pokemon';
set session my.generation = '1';
set session my.status = 'Dummy Text';
set session my.species = 'Dummy Text';
set session my.type_1 = 'Dummy Text';
set session my.type_2 = '';
set session my.height_m = '1.0';
set session my.weight_kg = '1.0';
set session my.ability_1 = '';
set session my.ability_2 = '';
set session my.ability_hidden = '';
set session my.total_points = '1';
set session my.hp = '1';
set session my.attack = '1';
set session my.defense = '1';
set session my.sp_attack = '1';
set session my.sp_defense = '1';
set session my.speed = '1';
set session my.catch_rate = '1';
set session my.base_friendship = '1';
set session my.base_experience = '1';
set session my.growth_rate = '';
set session my.egg_type_1 = '';
set session my.egg_type_2 = '';
set session my.percentage_male = '1.0';
set session my.percentage_female = '1.0';
set session my.egg_cycles = '1';
set session my.against_normal = '1.0';
set session my.against_fire = '1.0';
set session my.against_water = '1.0';
set session my.against_electric = '1.0';
set session my.against_grass = '1.0';
set session my.against_ice = '1.0';
set session my.against_fight = '1.0';
set session my.against_poison = '1.0';
set session my.against_ground = '1.0';
set session my.against_flying = '1.0';
set session my.against_psychic = '1.0';
set session my.against_bug = '1.0';
set session my.against_rock = '1.0';
set session my.against_ghost = '1.0';
set session my.against_dragon = '1.0';
set session my.against_dark = '1.0';
set session my.against_steel = '1.0';
set session my.against_fairy = '1.0';

insert into pokemon select * from current_setting('my.pokedex_number'),
                    current_setting('my.name'),
                    current_setting('my.generation'),
                    current_setting('my.status'),
                    current_setting('my.species'),
                    current_setting('my.type_1'),
                    current_setting('my.type_2'),
                    current_setting('my.height_m'),
                    current_setting('my.weight_kg'),
                    current_setting('my.ability_1'),
                    current_setting('my.ability_2'),
                    current_setting('my.ability_hidden'),
                    current_setting('my.total_points'),
                    current_setting('my.hp'),
                    current_setting('my.attack'),
                    current_setting('my.defense'),
                    current_setting('my.speed'),
                    current_setting('my.catch_rate'),
                    current_setting('my.base_friendship'),
                    current_setting('my.base_experience'),
                    current_setting('my.growth_rate'),
                    current_setting('my.egg_type_1'),
                    current_setting('my.egg_type_2'),
                    current_setting('my.percentage_male'),
                    current_setting('my.percentage_female'),
                    current_setting('my.egg_cycles'),
                    current_setting('my.against_normal'),
                    current_setting('my.against_fire'),
                    current_setting('my.against_water'),
                    current_setting('my.against_electric'),
                    current_setting('my.against_grass'),
                    current_setting('my.against_ice'),
                    current_setting('my.against_fight'),
                    current_setting('my.against_poison'),
                    current_setting('my.against_ground'),
                    current_setting('my.against_flying'),
                    current_setting('my.against_psychic'),
                    current_setting('my.against_bug'),
                    current_setting('my.against_rock'),
                    current_setting('my.against_ghost'),
                    current_setting('my.against_dragon'),
                    current_setting('my.against_dark'),
                    current_setting('my.against_steel'),
                    current_setting('my.against_fairy');