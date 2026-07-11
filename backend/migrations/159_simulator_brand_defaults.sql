-- Simulator fork branding defaults.
--
-- Only replace upstream defaults. If an operator already customized branding,
-- keep their value intact.

INSERT INTO settings (key, value, updated_at)
VALUES ('site_name', 'Simulator', CURRENT_TIMESTAMP)
ON CONFLICT (key) DO UPDATE
SET value = EXCLUDED.value,
    updated_at = CURRENT_TIMESTAMP
WHERE settings.value = '' OR settings.value = 'Sub2API';

INSERT INTO settings (key, value, updated_at)
VALUES ('site_subtitle', 'Build access for Sim Desktop and Sim Design', CURRENT_TIMESTAMP)
ON CONFLICT (key) DO UPDATE
SET value = EXCLUDED.value,
    updated_at = CURRENT_TIMESTAMP
WHERE settings.value = '' OR settings.value = 'Subscription to API Conversion Platform';

INSERT INTO settings (key, value, updated_at)
VALUES ('smtp_from_name', 'Simulator', CURRENT_TIMESTAMP)
ON CONFLICT (key) DO UPDATE
SET value = EXCLUDED.value,
    updated_at = CURRENT_TIMESTAMP
WHERE settings.value = '' OR settings.value = 'Sub2API';
