-- +goose Up
INSERT INTO songs (group_name, song_name, text, release_date, link)
VALUES
    ('Aurora', 'Dancing Lights',
     'The stars are falling, one by one\nThe sky is painted, the night begun\nI hear the whispers in the air\nA melody beyond compare\n\nCome and see the dancing lights\nThey shimmer, glow, and hold the night\nA symphony of dreams untold\nA wonderland of hues and gold\n\nEach spark ignites a fleeting flame\nA fleeting moment, yet untamed\nAnd though the dawn will take them soon\nThey''ll live forever in the moon',
     '2023-03-15', 'https://example.com/dancing-lights'),

    ('Nova', 'Endless Horizons',
     'Across the fields, the wind does call\nA voice that whispers to us all\nBeyond the hills, a secret lies\nWhere dreams are born, beneath the skies\n\nEndless horizons, wide and free\nA place for wanderers to be\nThe journey calls, we take our flight\nThrough day and stars, through endless night\n\nNo map to guide, no path to tread\nJust hope and courage, thoughts unsaid\nThe world awaits, so vast, so grand\nWith every step, we understand',
     '2022-11-20', 'https://example.com/endless-horizons'),

    ('Solace', 'Whispered Shadows',
     'In the quiet of the trees\nWhere shadows dance upon the leaves\nA voice so soft, it speaks of grace\nA fleeting echo, time can''t erase\n\nWhispered shadows, guiding me\nThrough the forest, silently\nThey tell of tales from years gone by\nOf love, of loss, of endless skies\n\nAnd in the stillness, I find peace\nA bond with nature, a sweet release\nFor every shadow holds a light\nA secret warmth within the night',
     '2021-08-11', 'https://example.com/whispered-shadows'),

    ('Eclipse', 'Fading Glow',
     'The sun descends beyond the sea\nA golden glow, a memory\nThe night prepares to claim its reign\nThe stars emerge, a soft refrain\n\nFading glow, the evening sighs\nA tender hush in twilight skies\nEach moment whispers, hold it near\nFor time will steal it, year by year\n\nBut even when the glow is gone\nIts warmth remains, it lingers on\nA fleeting beauty, pure and sweet\nIn fading glow, my soul competes',
     '2020-05-25', 'https://example.com/fading-glow'),

    ('Lumos', 'Silent Dawn',
     'Before the light, the earth does sleep\nIn shadows vast and silence deep\nA world awaits the first new ray\nTo cast the darkness far away\n\nSilent dawn, a moment still\nA quiet peace, a gentle thrill\nThe world awakens, breathes anew\nA canvas blank, awaiting hues\n\nWith every breath, a promise grows\nOf endless dreams, of paths unknown\nThe silent dawn, a fleeting spark\nThat chases shadows from the dark',
     '2019-12-30', 'https://example.com/silent-dawn'),

    ('Serenity', 'Crimson Sky',
     'A fiery blaze, the sky ignites\nA symphony of reds and whites\nThe day departs, the night arrives\nIn crimson hues, the world survives\n\nCrimson sky, a fleeting sight\nA dance of fire, a kiss of night\nIts beauty strikes the soul with awe\nA wonder no one ever saw\n\nAnd as it fades, its memory stays\nA painted dream of fleeting days\nThe crimson sky, a love untold\nA fiery passion, brave and bold',
     '2018-07-19', 'https://example.com/crimson-sky'),

    ('Echo', 'Timeless River',
     'A river flows, it knows no end\nThrough time and space, it twists and bends\nIt carries tales of ancient lands\nOf fleeting hearts and shifting sands\n\nTimeless river, gently sings\nOf broken hearts and mended wings\nIts song, a hymn to time itself\nA flowing treasure, nature''s wealth\n\nAnd though it journeys ever far\nIt carries with it every scar\nA timeless river, wild and free\nIt holds the world''s eternity',
     '2017-04-12', 'https://example.com/timeless-river'),

    ('Haven', 'Golden Shore',
     'Beneath the waves, the shore does gleam\nA treasure lost within a dream\nIts sands of gold, its skies so clear\nA promised land, so far, so near\n\nGolden shore, a beacon bright\nIt calls to wanderers through the night\nA safe embrace, a peaceful rest\nA place where hearts find what is best\n\nAnd though the journey may be long\nThe golden shore inspires song\nA hope that keeps the soul alive\nA goal for which we always strive',
     '2016-10-03', 'https://example.com/golden-shore'),

    ('Celeste', 'Silver Moon',
     'The silver moon, it lights the way\nThrough shadows vast, it holds its sway\nA gentle glow, so calm, serene\nA guide for all who seek unseen\n\nSilver moon, a friend at night\nIt comforts all with gentle light\nA silent guardian in the sky\nA watchful eye for those who cry\n\nAnd when the sun begins to rise\nThe silver moon bids soft goodbyes\nYet in its place, it leaves a trace\nA hope to light the darkest space',
     '2015-06-18', 'https://example.com/silver-moon'),

    ('Starlight', 'Echoes of Eternity',
     'The stars above, they sing a song\nOf endless time, of right and wrong\nEach echo carries through the night\nA message clear, a purest light\n\nEchoes of eternity\nThey speak of what will always be\nA story written in the sky\nOf dreams that never say goodbye\n\nAnd though the stars may fade away\nTheir echoes linger, here to stay\nA timeless truth, a mystery\nThe echoes of eternity',
     '2014-01-15', 'https://example.com/echoes-of-eternity');

-- +goose Down
DELETE FROM songs;
