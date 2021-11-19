CREATE TABLE IF NOT EXISTS public.jobs_items(
    id          SERIAL PRIMARY KEY,
    official_id INTEGER UNIQUE NOT NULL,
    gfx_id      INTEGER,
    title_fr    VARCHAR(64) NOT NULL,
    title_en    VARCHAR(64),
    title_es    VARCHAR(64),
    title_pt    VARCHAR(64)
);
