-- Rename the previous fork branding from TanStarter to Simulator.
--
-- This migration is intentionally narrow: it only updates values that were set
-- by the earlier fork branding migration, preserving operator customizations.

UPDATE settings
SET value = 'Simulator',
    updated_at = CURRENT_TIMESTAMP
WHERE key = 'site_name'
  AND value = 'TanStarter';

UPDATE settings
SET value = 'Simulator',
    updated_at = CURRENT_TIMESTAMP
WHERE key = 'smtp_from_name'
  AND value = 'TanStarter';

UPDATE settings
SET value = 'Build access for Sim Desktop and Sim Design',
    updated_at = CURRENT_TIMESTAMP
WHERE key = 'site_subtitle'
  AND value IN (
    'Subscription to API Conversion Platform',
    'Build access for Sim Desktop and Sim Design'
  );
